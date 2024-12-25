package fullcycle.balance_account_manager.repository;

import fullcycle.balance_account_manager.entity.AccountEntity;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface AccountRepository extends JpaRepository<AccountEntity, String> {
}
