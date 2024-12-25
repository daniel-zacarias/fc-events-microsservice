package fullcycle.balance_account_manager.event;


import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.UUID;

public class BalanceUpdatedEvent {

    @JsonProperty("Name")
    private String name;

    @JsonProperty("Payload")
    private Payload payload;

    public static class Payload {

        @JsonProperty("account_id_from")
        private String accountIdFrom;

        @JsonProperty("account_id_to")
        private String accountIdTo;

        @JsonProperty("balance_account_id_from")
        private double balanceAccountIdFrom;

        @JsonProperty("balance_account_id_to")
        private double balanceAccountIdTo;

        public String getAccountIdFrom() {
            return accountIdFrom;
        }

        public void setAccountIdFrom(String accountIdFrom) {
            this.accountIdFrom = accountIdFrom;
        }

        public String getAccountIdTo() {
            return accountIdTo;
        }

        public void setAccountIdTo(String accountIdTo) {
            this.accountIdTo = accountIdTo;
        }

        public double getBalanceAccountIdFrom() {
            return balanceAccountIdFrom;
        }

        public void setBalanceAccountIdFrom(double balanceAccountIdFrom) {
            this.balanceAccountIdFrom = balanceAccountIdFrom;
        }

        public double getBalanceAccountIdTo() {
            return balanceAccountIdTo;
        }

        public void setBalanceAccountIdTo(double balanceAccountIdTo) {
            this.balanceAccountIdTo = balanceAccountIdTo;
        }
    }

    // Getters and Setters for "name" and "payload"
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Payload getPayload() {
        return payload;
    }

    public void setPayload(Payload payload) {
        this.payload = payload;
    }
}

