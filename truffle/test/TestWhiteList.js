var SuperCoin = artifacts.require("./SuperCoin.sol");

contract('SuperCoin', function(accounts) {

    // 测试白名单功能
    it("should address on the white list", function() {
        var meta;

        return SuperCoin.deployed().then(function(instance) {
            meta = instance;

            // 验证新账号是否在白名单
            return meta.whitelist(accounts[1], { from: accounts[0] });
        }).then(function(exists) {
            assert.isFalse(exists, "account 1 shuld not on the white list ");
        }).then(function() {

            // 验证添加到白名单功能
            return meta.addToWhitelist(accounts[1], { from: accounts[0] });
        }).then(function(result) {
            console.log("addToWhitelist successful!" + result.tx);
        }).then(function() {

            // 再次验证帐号是否在白名单
            return meta.whitelist(accounts[1], { from: accounts[0] });
        }).then(function(exists) {
            assert.isTrue(exists, "account 1 shuld on the white list ");
        });
    });
});