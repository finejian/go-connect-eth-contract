var SuperCoin = artifacts.require("./SuperCoin.sol");

contract('SuperCoin', function(accounts) {

    // 测试铸币权限，必须是 owner 给白名单铸币
    it("should mint coin correctly", function() {
        var meta;

        return SuperCoin.deployed().then(function(instance) {
            meta = instance;

            // 查询account2余额
            return meta.balanceOf(accounts[2]);
        }).then(function(balance) {
            var num = new Number(balance.valueOf())
            console.log("account2 balance: %s", num.toLocaleString());

            // 添加到白名单
            return meta.addToWhitelist(accounts[2], { from: accounts[0] });
        }).then(function() {
            // 验证白名单地址铸币

            var targetAccount = accounts[2]
            return meta.mint(targetAccounts, 102400, { from: accounts[0] });
        }).then(function(success) {

            // 查询account2余额
            return meta.balanceOf(accounts[2]);
        }).then(function(balance) {
            var num = new Number(balance.valueOf())
            console.log("account2 balance: %s", num.toLocaleString());
        });
    });
});