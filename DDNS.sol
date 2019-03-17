pragma solidity >=0.4.22 <0.6.0;

import "./common/Ownable.sol";
import "./libs/SafeMath.sol";

contract DDNS is Ownable {
    using SafeMath for uint256;

    constructor() public {
        register("google", "com", "9.9.9.9");
    }

    /** === STRUCTS START === */

    // struct to hold the domain record
    struct DomainRecord {
        bytes name;
        bytes12 tld;
        address owner;
        bytes15[] ipAddress;
        uint expires;
    }

    // struct to hold the receipt to track payments made
    struct Receipt {
        uint amountPaid; // in wei
        uint timestamp;
        uint expires;
    }

    /** === STRUCTS END === */

    /** === CONTRACT VARIABLES AND CONSTANTS START === */

    // Constants
    uint constant public DOMAIN_NAME_COST = 1 ether;
    uint constant public DOMAIN_NAME_COST_SHORT_ADDITION = 1 ether;
    uint constant public DOMAIN_EXPIRATION_DATE = 365 days;
    uint8 constant public DOMAIN_NAME_MIN_LENGTH = 5;
    uint8 constant public DOMAIN_NAME_EXPENSIVE_LENGTH = 8;
    uint8 constant public TLD_MIN_LENGTH = 1;
    bytes1 constant public BYTES_DEFAULT_VALUE = bytes1(0x00);

    // Contract Variables
    mapping (bytes32 => DomainRecord) public domainNames; // mapping domain record to a domain hash
    mapping (address => bytes32[]) public paymentReceipts; // mapping all the receipt hashes to an addres
    mapping (bytes32 => Receipt) public receiptDetails; // mapping a receipt to the receipt Hash

    /** === CONTRACT VARIABLES AND CONSTANTS END === */

    /** === MODIFIERS START */

    // check if domain is available
    modifier isAvailable(bytes memory domain, bytes12 tld) {
        bytes32 domainHash = getDomainHash(domain, tld);
        require(
            domainNames[domainHash].expires < block.timestamp,
            "Domain Name is not available"
        );
        _;
    }

    // check if initiator is owner of domain when required
    modifier isDomainOwner(bytes memory domain, bytes12 tld) {
        bytes32 domainHash = getDomainHash(domain, tld);
        require(
            domainNames[domainHash].owner == msg.sender,
            "Transaction initiator is not domain owner"
        );
        _;
    }

    // // check if amount paid sufficient to aquire domain
    // modifier verifyDomainNamePayment(bytes memory domain) {
    //     uint domainPrice = getPrice(domain);
    //     require(
    //         msg.value >= domainPrice,
    //         "Insufficient amount"
    //     );
    //     _;
    // }

    // // check if domain length is greater than minimum length
    // modifier isDomainLengthValid(bytes memory domain) {
    //     require(
    //         domain.length >= DOMAIN_NAME_MIN_LENGTH,
    //         "Domain name is too short"
    //     );
    //     _;
    // }

    // // check if required TLD is of valid length
    // modifier isTldAllowed(bytes12 tld) {
    //     require(
    //         tld.length >= TLD_MIN_LENGTH,
    //         "Required TLD is too short"
    //     );
    //     _;
    // }

    /** === MODIFIERS END === */

    /** === EVENTS START === */

    // event to signify registration of a new domain
    event DomainNameRegistered(
        uint indexed timestamp,
        bytes domainName, 
        bytes12 tld
    );

    // event to signify renewal of an expired/expiring domain
    event DomainNameRenewed(
        uint indexed timestamp,
        bytes domainName,
        bytes12 tld,
        address indexed owner
    );

    // event to signify a change in the IP of a domain
    event DomainNameEdited(
        uint indexed timestamp,
        bytes domainName,
        bytes12 tld,
        bytes15 newIp
    );

    // event to signify transfer of a domain to a new owner
    event DomainNameTransferred(
        uint indexed timestamp,
        bytes domainName,
        bytes12 tld,
        address indexed owner,
        address newOwner
    );

    // event to signify returning of change to transaction initiator
    event PurchaseChangeReturned(
        uint indexed timestamp,
        address indexed _owner,
        uint amount
    );

    // event to signify the saving of a receipt
    event ReceiptSaved(
        uint indexed timestamp,
        bytes domainName,
        uint amountInWei,
        uint expires
    );

    /** === EVENTS END === */

    /** === FUNCTIONS START === */

    function getPrice(bytes memory domain) public pure returns (uint){
        if(domain.length < DOMAIN_NAME_EXPENSIVE_LENGTH) {
            return DOMAIN_NAME_COST + DOMAIN_NAME_COST_SHORT_ADDITION;
        }

        return DOMAIN_NAME_COST;
    }

    function getDomainHash(bytes memory domain, bytes12 topLevel) public pure returns(bytes32){
        return keccak256(abi.encodePacked(domain, topLevel));
    }

    function register(bytes memory domain, bytes12 tld, bytes15 ip) public payable
        // isDomainLengthValid(domain)
        isAvailable(domain, tld){
            bytes32 domainHash = getDomainHash(domain, tld);
            bytes15[] memory ipAddresses = new bytes15[](1);
            ipAddresses[0] = ip;
            DomainRecord memory newDomain = DomainRecord({
                name: domain,
                tld: tld,
                owner: msg.sender,
                ipAddress: ipAddresses,
                expires: block.timestamp + DOMAIN_EXPIRATION_DATE
            });
            domainNames[domainHash] = newDomain;

            emit DomainNameRegistered(
                block.timestamp,
                domain, 
                tld
            );
    }

    function edit(bytes memory domain, bytes12 tld, bytes15 newIp) public isDomainOwner(domain, tld) {
        bytes32 domainHash = getDomainHash(domain, tld);
        domainNames[domainHash].ipAddress.push(newIp);
        emit DomainNameEdited(block.timestamp, domain, tld, newIp);
    }

    function getIp(bytes memory domain, bytes12 tld) public view returns (bytes15[] memory outArr) {
        bytes32 domainHash = getDomainHash(domain, tld);
        // bytes15[] outArr;
        // return domainNames[domainHash].ipAddress;
        outArr = domainNames[domainHash].ipAddress;
    }

    /** === FUNCTIONS END === */
}