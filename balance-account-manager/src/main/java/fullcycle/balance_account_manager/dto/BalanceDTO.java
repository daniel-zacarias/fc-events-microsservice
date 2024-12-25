package fullcycle.balance_account_manager.dto;

public class BalanceDTO {
    private Double balance;

    public BalanceDTO(Double balance) {
        this.balance = balance;
    }

    public Double getBalance() {
        return balance;
    }

    public void setBalance(Double balance) {
        this.balance = balance;
    }
}
