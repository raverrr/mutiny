package main

import (
        "fmt"
        "strings"
        "github.com/fatih/color"
        "bufio"
        "context"
        "flag"
        "log"
        "os"
        "sync"
        "time"
        "github.com/chromedp/cdproto/page"
        "github.com/chromedp/chromedp"
)

var res2 string
var output string
var concurrency int
var customCookies string
var rateLimit int

func displayBanner() {
        fmt.Println("")
         color.Blue("-------------------------------------------------------------------")
        fmt.Println("")
          color.Red("               █▀▄▀█   ▄     ▄▄▄▄▀ ▄█    ▄  ▀▄    ▄ ")
          color.Red("               █ █ █    █ ▀▀▀ █    ██     █   █  █  ")
          color.Red("               █ ▄ █ █   █    █    ██ ██   █   ▀█   ")
          color.Red("               █   █ █   █   █     ▐█ █ █  █   █    ")
          color.Red("                  █  █▄ ▄█  ▀       ▐ █  █ █ ▄▀     ")
          color.Red("                 ▀    ▀▀▀             █   ██        ")
        fmt.Println("                                               ")
        fmt.Println("                                                                   ")
         color.Blue("                      Mutiny Version: 1.0.1                   ")
        fmt.Println("                                                                   ")
          color.Red("-- You're off the edge of the app, mate. Here there be monsters. --")
        fmt.Println("                                                                   ")
         color.Blue("-------------------------------------------------------------------")
        fmt.Println("")
}


func main() {
        displayBanner()
        log.SetFlags(0)
        flag.StringVar(&output, "o", "", "Specify the output file where results will be saved.")
        flag.IntVar(&concurrency, "c", 5, "Set the number of concurrent threads for processing.")
        flag.IntVar(&rateLimit, "r", 0, "Set a rate limit for requests in milliseconds.")
        flag.StringVar(&customCookies, "cookies", "", "Provide custom cookies for authentication while making requests.")
        flag.Parse()

        jsSnippet1 := `let functionCallsArray = [];
const sensitiveDataList = Array.from(new URLSearchParams(window.location.search).values()).filter(value => value.length > 3);
let functionCallsCount = {};
function logFunctionCall(message) {
  if (functionCallsCount[message]) {
    functionCallsCount[message]++;
  } else {
    functionCallsCount[message] = 1;
  }
}

function generateSummary() {
  let summary = '';
  const sortedEntries = Object.entries(functionCallsCount).sort((a, b) => a[0].localeCompare(b[0]));
  for (const [message, count] of sortedEntries) {
    summary += '\n---' + message + ' ' + count + ' time(s)';
  }
  return summary;
}

function wrapInnerHTML() {
  const originalInnerHTML = Object.getOwnPropertyDescriptor(Element.prototype, 'innerHTML').set;
  Object.defineProperty(Element.prototype, 'innerHTML', {
    set: function(value) {
      if (value.includes('<script>')) {
        logFunctionCall('innerHTML: Script tag insertion');
      }
      return originalInnerHTML.call(this, value);
    }
  });
}

function wrapDocumentWrite() {
  const originalDocumentWrite = document.write;
  document.write = function(string) {
    if (string.includes('<script>')) {
      logFunctionCall('document.write: Script tag insertion');
    }
    return originalDocumentWrite.apply(this, arguments);
  };
}

function wrapSetAttribute() {
  const originalSetAttribute = Element.prototype.setAttribute;
  Element.prototype.setAttribute = function(name, value) {
    if (name.startsWith('on')) {
      logFunctionCall('setAttribute: Event handler ' + name);
    }
    return originalSetAttribute.apply(this, arguments);
  };
}

function wrapWindowOpen() {
  const originalWindowOpen = window.open;
  window.open = function(url) {
    const currentOrigin = window.location.origin;
    const newUrl = new URL(url, currentOrigin);
    if (newUrl.origin !== currentOrigin) {
      logFunctionCall('window.open: Different origin');
    }
    return originalWindowOpen.apply(this, arguments);
  };
}

function wrapInsertAdjacentHTML() {
  const originalInsertAdjacentHTML = Element.prototype.insertAdjacentHTML;
  Element.prototype.insertAdjacentHTML = function(position, text) {
    if (text.includes('<script>')) {
      logFunctionCall('insertAdjacentHTML: Script tag insertion');
    }
    return originalInsertAdjacentHTML.apply(this, arguments);
  };
}

function wrapAjaxSend() {
  const originalSend = XMLHttpRequest.prototype.send;
  XMLHttpRequest.prototype.send = function(body) {
    const hasSensitiveData = sensitiveDataList.some(data => (this.responseURL && this.responseURL.includes(data)));
    if (hasSensitiveData) {
      logFunctionCall('AJAX: Data from page URL in AJAX request - Method: ' + this.method + ', URL: ' + this.responseURL);
    }
    return originalSend.apply(this, arguments);
  };
}

function wrapFetch() {
  const originalFetch = window.fetch;
  window.fetch = function(input, init) {
    const url = typeof input === 'string' ? input : input.url;
    logFunctionCall('Fetch: ' + url);
    return originalFetch.apply(this, arguments);
  };
}

function wrapEval() {
  const originalEval = window.eval;
  window.eval = function(script) {
    logFunctionCall('Eval called');
    return originalEval.apply(this, arguments);
  };
}

function wrapFormSubmit() {
  HTMLFormElement.prototype.realSubmit = HTMLFormElement.prototype.submit;
  HTMLFormElement.prototype.submit = function() {
    logFunctionCall('Form submitted');
    this.realSubmit();
  };
}

function wrapLocalStorage() {
  const originalSetItem = Storage.prototype.setItem;
  Storage.prototype.setItem = function(key, value) {
    logFunctionCall('LocalStorage: setItem called');
    return originalSetItem.apply(this, arguments);
  };
}

function wrapSessionStorage() {
  const originalSessionSetItem = sessionStorage.setItem;
  sessionStorage.setItem = function(key, value) {
    logFunctionCall('SessionStorage: setItem called');
    return originalSessionSetItem.apply(this, arguments);
  };
}

function wrapSendBeacon() {
  const originalSendBeacon = Navigator.prototype.sendBeacon;
  Navigator.prototype.sendBeacon = function(url, data) {
    logFunctionCall('sendBeacon called');
    return originalSendBeacon.apply(this, arguments);
  };
}

function wrapWebSocket() {
  const originalWebSocket = window.WebSocket;
  window.WebSocket = function(url, protocols) {
    logFunctionCall('WebSocket: ' + url);
    return new originalWebSocket(url, protocols);
  };
}

function wrapCreateElement() {
  const originalCreateElement = document.createElement;
  document.createElement = function(tagName, options) {
    logFunctionCall('createElement: ' + tagName);
    return originalCreateElement.apply(this, arguments);
  };
}

function wrapAppendChild() {
  const originalAppendChild = Node.prototype.appendChild;
  Node.prototype.appendChild = function(child) {
    logFunctionCall('appendChild called');
    return originalAppendChild.apply(this, arguments);
  };
}

// Conditionally wrap jQuery
if (typeof $ !== "undefined") {
  function wrapJQueryAjax() {
    const originalJQueryAjax = $.ajax;
    $.ajax = function(options) {
      logFunctionCall('jQuery.ajax: ' + options.url);
      return originalJQueryAjax.apply(this, arguments);
    };
  }
  wrapJQueryAjax();
}
function wrapJQueryAjax() {
  const originalJQueryAjax = $.ajax;
  $.ajax = function(options) {
    logFunctionCall('jQuery.ajax: ' + options.url);
    return originalJQueryAjax.apply(this, arguments);
  };
}

function wrapHistoryAPI() {
  const originalPushState = History.prototype.pushState;
  History.prototype.pushState = function(state, title, url) {
    logFunctionCall('History API: pushState');
    return originalPushState.apply(this, arguments);
  };
}

function wrapCloneNode() {
  const originalCloneNode = Element.prototype.cloneNode;
  Element.prototype.cloneNode = function(deep) {
    logFunctionCall('cloneNode called');
    return originalCloneNode.apply(this, arguments);
  };
}

function wrapRemoveChild() {
  const originalRemoveChild = Node.prototype.removeChild;
  Node.prototype.removeChild = function(child) {
    logFunctionCall('removeChild called');
    return originalRemoveChild.apply(this, arguments);
  };
}

function wrapGetCurrentPosition() {
  const originalGetCurrentPosition = navigator.geolocation.getCurrentPosition;
  navigator.geolocation.getCurrentPosition = function(successCallback, errorCallback, options) {
    logFunctionCall('getCurrentPosition called');
    return originalGetCurrentPosition.apply(this, arguments);
  };
}

function wrapAddEventListener() {
  const originalAddEventListener = EventTarget.prototype.addEventListener;
  EventTarget.prototype.addEventListener = function(type, listener, options) {
    logFunctionCall('addEventListener: ' + type);
    return originalAddEventListener.apply(this, arguments);
  };
}

function wrapPostMessage() {
  const originalPostMessage = window.postMessage;
  window.postMessage = function(message, targetOrigin, transfer) {
    logFunctionCall('postMessage: ' + targetOrigin);
    return originalPostMessage.apply(this, arguments);
  };
}

function wrapFetchAbort() {
  const originalAbort = AbortController.prototype.abort;
  AbortController.prototype.abort = function() {
    if (this.signal && this.signal.aborted && this.signal.controller) {
      const request = this.signal.controller.request;
      const url = request ? request.url : "Unknown URL";
      logFunctionCall('Fetch request aborted: ' + url);
    } else {
      logFunctionCall('Fetch request aborted: Unknown URL');
    }
    return originalAbort.apply(this, arguments);
  };
}

function wrapSetInterval() {
  const originalSetInterval = window.setInterval;
  window.setInterval = function(callback, delay) {
    logFunctionCall('setInterval called');
    return originalSetInterval.apply(this, arguments);
  };
}

function wrapSetTimeout() {
  const originalSetTimeout = window.setTimeout;
  window.setTimeout = function(callback, delay) {
    logFunctionCall('setTimeout called');
    return originalSetTimeout.apply(this, arguments);
  };
}


function wrapAllFunctions() {
  wrapInnerHTML();
  wrapDocumentWrite();
  wrapSetAttribute();
  wrapWindowOpen();
  wrapInsertAdjacentHTML();
  wrapAjaxSend();
  wrapFetch();
  wrapEval();
  wrapFormSubmit();
  wrapLocalStorage();
  wrapSessionStorage();
  wrapSendBeacon();
  wrapWebSocket();
  wrapCreateElement();
  wrapAppendChild();
//  wrapJQueryAjax(); breaks. MAde conditional
  wrapHistoryAPI();
  wrapCloneNode();
  wrapRemoveChild();
  wrapGetCurrentPosition();
  wrapAddEventListener();
  wrapPostMessage();
  wrapFetchAbort();
  wrapSetInterval();
  wrapSetTimeout();
}

wrapAllFunctions();`

        jsSnippet2 := `generateSummary();`

        var wg sync.WaitGroup
        jobs := make(chan string)

        for i := 0; i < concurrency; i++ {
                wg.Add(1)
                go func() {
                        for requestURL := range jobs {
                                if rateLimit > 0 {
                                        time.Sleep(time.Duration(rateLimit) * time.Millisecond)
                                }

                                ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
                                defer cancel()

                                copts := append(chromedp.DefaultExecAllocatorOptions[:],
                                        chromedp.Flag("ignore-certificate-errors", true),
                                )

                                ctx, cancel = chromedp.NewExecAllocator(ctx, copts...)
                                defer cancel()

                                ctx, cancel = chromedp.NewContext(ctx)
                                defer cancel()

                                tasks := chromedp.Tasks{
                                        chromedp.ActionFunc(func(ctx context.Context) error {
                                                if customCookies != "" {
                                                        expr := `document.cookie = "` + customCookies + `";`
                                                        return chromedp.Evaluate(expr, nil).Do(ctx)
                                                }
                                                return nil
                                        }),
                                        chromedp.ActionFunc(func(ctx context.Context) error {
                                                _, _, _, err := page.Navigate(requestURL).Do(ctx)
                                                return err
                                        }),
                                        chromedp.Evaluate(jsSnippet1, nil),
                                        chromedp.WaitReady("body"),
                                        chromedp.Sleep(2*time.Second),
                                        chromedp.Evaluate(jsSnippet2, &res2),
                                }
                                  //This will do, fix later.
                                if err := chromedp.Run(ctx, tasks); err != nil {
                                 if !strings.Contains(err.Error(), "generateSummary is not defined") {
                                  log.Printf("[-] %s: %s", requestURL, err.Error())
                                 }
                                } else {
                                    if res2 != "" {
                                     color.Blue("[+] %s:", requestURL)
                                     color.Green("%s", res2)
                                     log.Println()

                                     if output != "" {
                                      err := writeToFile(output, fmt.Sprintf("\n[+] %s:\n%s\n", requestURL, res2))
                                      if err != nil {
                                       log.Printf("[-] Failed to write to output file: %s", err.Error())
                                       }
                                     }
                                   }
                                 }
                                 cancel()

                        }
                        wg.Done()
                }()
        }

        sc := bufio.NewScanner(os.Stdin)
        for sc.Scan() {
                jobs <- sc.Text()
        }
        close(jobs)
        wg.Wait()
}

func writeToFile(filename, content string) error {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()
    _, err = f.WriteString(content)
    return err
}
