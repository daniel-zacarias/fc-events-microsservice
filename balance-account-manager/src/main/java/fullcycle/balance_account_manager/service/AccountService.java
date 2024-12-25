package fullcycle.balance_account_manager.service;

import fullcycle.balance_account_manager.entity.AccountEntity;
import fullcycle.balance_account_manager.repository.AccountRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
public class AccountService {
    @Autowired
    private AccountRepository accountRepository;

    public Double getBalanceByAccountId(String id){
        AccountEntity account = accountRepository.findById(id).orElseThrow(() -> new RuntimeException("Conta não encontrada!"));
        return account.getBalance();
    }
}
