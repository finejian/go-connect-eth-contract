var SuperCoin = artifacts.require("./SuperCoin.sol");

contract('SuperCoin', function(accounts) {

    // 测试转币
    it("should send coin correctly", function() {
        var meta;

        // Get initial balances of first and second account.
        var account_one = accounts[0];
        var account_two = accounts[1];

        var account_one_starting_balance;
        var account_two_starting_balance;
        var account_one_ending_balance;
        var account_two_ending_balance;

        var amount = 1050000000 * (10 ** 18);

        return SuperCoin.deployed().then(function(instance) {
            meta = instance;
            // 拿到帐号1的初始余额
            return meta.balanceOf(account_one);
        }).then(function(balance) {
            account_one_starting_balance = balance.toNumber();

            // 拿到帐号2的初始余额
            return meta.balanceOf(account_two);
        }).then(function(balance) {
            account_two_starting_balance = balance.toNumber();

            // 帐号1给帐号2转走1050000000
            return meta.transfer(account_two, amount, { from: account_one });
        }).then(function() {
            return meta.balanceOf(account_one);
        }).then(function(balance) {
            account_one_ending_balance = balance.toNumber();
            return meta.balanceOf(account_two);
        }).then(function(balance) {
            account_two_ending_balance = balance.toNumber();

            assert.equal(account_one_starting_balance, amount, "account_one_starting_balance: " + account_one_starting_balance);
            assert.equal(account_two_starting_balance, 0, "account_two_starting_balance: " + account_two_starting_balance);

            assert.equal(account_one_ending_balance, 0, "account_one_ending_balance: " + account_one_ending_balance);
            assert.equal(account_two_ending_balance, amount, "account_two_ending_balance: " + account_two_ending_balance);
        });
    });
});