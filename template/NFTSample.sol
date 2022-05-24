// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//imports a specific Solidity contract from the Open Zeppelin Github repository; 
//check the latest version # and update it in the import statement's path as needed
/* 
  import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.3.0/contracts/token/ERC1155/ERC1155.sol";
  import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.3.0/contracts/access/Ownable.sol";
 */
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


/**
* @author Forrest Colyer (AWS), inherits from Open Zeppelin contracts
* @title A simple ERC1155 fungibility-agnostic token contract w/ access control and batch functions
*/
contract NFTSample is ERC1155, Ownable {
    
    //maps addresses to boolean flags for allowance and revocation of allowance to receive tokens
    //an address mapped to 'false' will be blocked from being sent tokens
    mapping(address => bool) internal allowedAddresses;

    //作品集名称
    string public name;

    //作品集描述
    string public description;

    address public txorigin;

    address public msgsender;

    constructor (string memory contract_name, string memory contract_description, string memory contract_uri, address contract_owner) ERC1155(contract_uri) 
    {
        name = contract_name;
        description = contract_description;
        setURI(contract_uri);
        msgsender = msg.sender;
        txorigin = txorigin;
        super._transferOwnership(contract_owner);
    }

    /*
    used to change metadata, only owner access
    https://www.quicknode.com/guides/solidity/how-to-create-and-deploy-a-factory-erc-1155-contract-on-polygon-using-truffle
    */
    function setURI(string memory newuri) public onlyOwner {
        super._setURI(newuri);
    }

    function mint(address to, uint256 id, uint256 amount, bytes memory data) public onlyOwner {
        super._mint(to, id, amount, data);
    }
    
    function mintBatch(address to, uint256[] memory ids, uint256[] memory amounts, bytes memory data) public onlyOwner {
        super._mintBatch(to, ids, amounts, data);
    }
    
    function burnBatch(address to, uint256[] memory ids, uint256[] memory amounts) public onlyOwner {
        super._burnBatch(to, ids, amounts);
    }
    
    
    /**
     * allowAddress: adds an Ethereum address to the allow list to transfer tokens
     * @param from - the Ethereum address of the sender to add to the allow list
     * @param allow - the boolean value denoting whether the to address is allowed to transfer/send tokens
     * 
     * @dev provide the address for the sender and boolean 'false' to deny and 'true' to allow token transfers
     */
    function allowAddress(address from, bool allow) public onlyOwner {
        allowedAddresses[from] = allow;
    }
    
    
    function _beforeTokenTransfer(address operator, address from, address to, uint256[] memory ids, uint256[] memory amounts, bytes memory data)
        internal virtual override
    {
        super._beforeTokenTransfer(operator, from, to, ids, amounts, data);

        if(msg.sender != owner()) { 
            require(_validSender(from), "ERC155WithSafeTransfer: this address does not have permission to transfer tokens");
        }
    }

    /**
     * _validSender: checks the 
     * @param from - the Ethereum address of the sender to check against the allow list
     * 
     * @dev the allow list is only checked for 'true' or 'false'; transfers initiated by the contract owner are not checked
     * 
     */
    function _validSender(address from) private view returns (bool) {
        //add logic for 'magic phrase' here to validate recipients?
        return allowedAddresses[from];
        
    }
}
