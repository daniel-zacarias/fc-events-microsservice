package fullcycle.balance_account_manager.controller;

import fullcycle.balance_account_manager.dto.BalanceDTO;
import fullcycle.balance_account_manager.service.AccountService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
public class BalanceController {

    @Autowired
    private AccountService accountService;

    @GetMapping("/balance/{account_id}")
    public ResponseEntity<BalanceDTO> one(@PathVariable String account_id) {
        Double accountBalance = accountService.getBalanceByAccountId(account_id);
        BalanceDTO balanceDTO = new BalanceDTO(accountBalance);
        return ResponseEntity.ok(balanceDTO);
    }
}
