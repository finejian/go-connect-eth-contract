var Migrations = artifacts.require("./SuperCoin.sol");

module.exports = function(deployer) {
    deployer.deploy(Migrations);
};