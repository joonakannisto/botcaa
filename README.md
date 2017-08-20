# botcaa
Automatically generate CAA records from Certbot config

Certificate enrollment using domain name validation (DV) is subject to attacks on DNS as well as routing attacks such as BGP hijacking. 
DNSSEC helps to mitigate the DNS based attacks, but even if the DNS is secured 
the target IP addresses can be hijacked or MITMed (other attacks against routing/forwarding). 


CAA records[1] can be used to limit the CAs, which can issue certificates for a domain. 
This alone helps with the DV issue, as a BGP hijacker has to target a specific provider instead of any DV provider. 
Moreover, the specification allows to bind to an account-uri, which means that renewal requests can be made only with some specific account credentials. These constraints force the attacker to either attack the DNSSEC enabled nameserver and its keys, or the server that has the DV account credentials.  

This tool generates account-uri constrained CAA records from existing certbot renewal configuration. The tool parses the configuration, and outputs a CAA record line, which can be used in up to date versions of BIND. 

Note 2017-08-20: I haven't tested the tool on an actual nameserver yet, just that it produces correct looking output.   

[1] https://tools.ietf.org/html/draft-ietf-acme-caa
