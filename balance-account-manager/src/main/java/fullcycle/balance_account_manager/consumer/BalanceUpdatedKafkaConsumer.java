package fullcycle.balance_account_manager.consumer;

import com.fasterxml.jackson.databind.ObjectMapper;
import fullcycle.balance_account_manager.entity.AccountEntity;
import fullcycle.balance_account_manager.event.BalanceUpdatedEvent;
import fullcycle.balance_account_manager.repository.AccountRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

@Component
public class BalanceUpdatedKafkaConsumer {

    @Autowired
    private AccountRepository accountRepository;


    @KafkaListener(topics = "balances", groupId = "wallet")
    public void consumeBalanceUpdatedEvent(String message) {
        try {

            ObjectMapper objectMapper = new ObjectMapper();
            BalanceUpdatedEvent event = objectMapper.readValue(message, BalanceUpdatedEvent.class);


            BalanceUpdatedEvent.Payload payload = event.getPayload();

            AccountEntity accountFrom = accountRepository.findById(payload.getAccountIdFrom())
                    .orElseThrow(() -> new RuntimeException("Conta de origem não encontrada!"));
            AccountEntity accountTo = accountRepository.findById(payload.getAccountIdTo())
                    .orElseThrow(() -> new RuntimeException("Conta de destino não encontrada!"));

            accountFrom.setBalance(payload.getBalanceAccountIdFrom());
            accountTo.setBalance(payload.getBalanceAccountIdTo());


            accountRepository.save(accountFrom);
            accountRepository.save(accountTo);

            System.out.println("Saldos atualizados com sucesso. Conta de origem: " + accountFrom.getId() +
                    ", Conta de destino: " + accountTo.getId());
        } catch (Exception e) {
            System.err.println("Erro ao processar evento: " + e.getMessage());
        }
    }

}
