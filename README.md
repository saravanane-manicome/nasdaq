## Nasdaq Watcher

The Nasdaq Stock Market (/ˈnæzdæk/; National Association of Securities Dealers Automated Quotations Stock Market) is an 
American stock exchange based in New York City. It is the most active stock trading venue in the US by volume, and 
ranked second on the list of stock exchanges by market capitalization of shares traded, behind the New York Stock 
Exchange. The exchange platform is owned by Nasdaq, Inc., which also owns the Nasdaq Nordic stock market network 
and several U.S.-based stock and options exchanges.
Source: Wikipedia

This project serves as a hands-on for Go, Gin and gRPC

It is divided in three main parts which are:
* the provider: as the name implies it provides quotes for requested symbol
* the REST API: exposes a REST endpoint to request for a Nasdaq quote
* the watcher: periodically requests a quote for the given symbol and frequency of updates

The watcher sends requests to the REST API, which in turn sends a request to the provider
Each of these parts are microservices
