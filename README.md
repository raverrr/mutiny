# Mutiny
A tool for security research and bug bounty that uses a headless browser to wrap common JS functions and report on their usage as a page is loading. 
##
<img src="https://i.imgur.com/QA7eNGc.png" width="600" height="650">


# Wrapped functions:
  * innerHTML
  * documentWrite
  * setAttribute
  * windowOpen
  * insertAdjacentHTML
  * ajaxSend
  * fetch
  * eval
  * formSubmit
  * localStorage
  * sessionStorage
  * sendBeacon
  * webSocket
  * createElement
  * appendChild
  * JQueryAjax
  * historyAPI
  * cloneNode
  * removeChild
  * getCurrentPosition (Geolocation)
  * addEventListener
  * postMessage
  * abort (fetch abort)
  * setInterval
  * setTimeout


# Install
`go install github.com/raverrr/mutiny@latest`

# Usage
`cat urls.txt | mutiny`

# Help
` -c [int] Set the number of concurrent threads for processing. (default 5)          `

` -cookies [string] Provide custom cookies for authentication while making requests.`

`  -o [string] Specify the output file where results will be saved.                  `

`  -r [int] Set a rate limit for requests in milliseconds.                           `


[![Buy Me A Coffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/r4v3rrr)
