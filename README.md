Complete
=========

Complete is a general purpose completion handler system that is mobile compatible
with gomobile. Clients can follow the `Completionable` interface to write their
own completion handlers in their native language which will be executed upon
completion in Go.

Complete also provides a generic `Result` struct that can hold any kind of Go
data structure, even if it is not supported by the native language. It's up
to the Go library or program to provide type assertion methods to convert the
generic `Result` object into its actual type.

## Usage

### iOS
Because functions cannot be passed directly to go, each target language must
implement their own completion handler class that implements the 
`complete.Completionable` interface/protocol.

Here is an example completion handler in Swift:

```swift
class CompletionHandler: GoCompleteCompletionableProtocol {
    
    var success: ((result: GoCompleteResult!) -> Void)!
    var failure: ((msg: String!, error: NSError!) -> Void)!

    @objc func onSuccess(p0: GoCompleteResult!) {
        self.success(result: p0)
    }
    
    @objc func onFailure(p0: String!, p1: NSError!) {
        self.failure(msg: p0, error: p1)
    }
}

func onSuccess(success: ((result: GoCompleteResult!) -> Void)!, failure: ((msg: String!, error: NSError!) -> Void)!) -> CompletionHandler {
    let handler = CompletionHandler()
    handler.success = success
    handler.failure = failure
    return handler
}
```

#### Login example
Here is an example in Swift of logging inwith a
completion handler:

```swift
        // Login
        Login("picard", "123456", false, onSuccess({ (result) in
            print("Logged in!")
            }, failure: { (msg, error) in
                print("Failed to log in: \(msg)")
        }))
```

### Android
Here is an example completion handler in Java. There may be better way to
implement your completion handler:

```java
public class CompletionHandler implements Handlers.Completionable {

    private Success success;
    private Failure failure;

    public void OnSuccess(Result result) {
        this.success.call(result);
    }
    public void OnFailure(String error) {
        this.failure.call(error);
    }

    public static CompletionHandler onSuccess(Success success, Failure failure) {
        CompletionHandler handler = new CompletionHandler();
        handler.success = success;
        handler.failure = failure;

        return handler;
    }
}

// Used for calling the success handler.
interface Success {
    void call(Result result);
}

// Used for calling the failure handler.
interface Failure {
    void call(String error);
}
```

#### Login example
Here is an example in Java of logging in with a
completion handler:

```java
// Login
Login("picard", "123456", false,
    CompletionHandler.onSuccess(
        new Success() {
            @Override
            public void call(Complete.Result result) {
                System.out.println("Logged in!");
            }
        },
        new Failure() {
            @Override
            public void call(String error) {
                System.out.println(error);
            }
        }
    )
);
```
