# gotick

A simple go package to run periodic task

## Usage
    
1. Import package

    ```        
    import "gotick/goptick"
    ```

2. Define

    ```        
    // Duration
    d := 1 * time.Second

    // Function
    func foo(q chan bool){

    }

    ```
 
3. Invoke

    ```
    tic.Gotick(1*d, foo)
    ```