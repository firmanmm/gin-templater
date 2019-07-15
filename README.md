# Gin Templater
A specialized gin gonic framework's templating system. It allows you to use multi level folders for the sake of your convenience and provide you with the ability to Hot Reload your HTML Code without restarting your Gin Web Server. Useful if you don't like to wait to recompile your Gin instance and get your web interface response imediately.

## Usage
By default Gin Templater will read from "view" directory and will output it's build to "cache/view" but can be overriden in the config. You can also disable the Hot Reload feature in the directory.

Here is the directory structure for the view : 
```
view/
    -home/
        -deep.html
    -part/
        -greetings.html
        -who.html
    -home.html
```
And here is the minimum working code to use Gin Templater :
```
engine := gin.Default()
config := gintemplater.NewConfig()
templater := gintemplater.NewTemplater(engine, config)
templater.Run()

engine.GET("/", func(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "home.html", gin.H{
        "message": "Hello!",
        "who":     "A Message",
    })
})

engine.GET("/deep", func(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "home.deep.html", gin.H{
        "message": "Hello!",
        "who":     "A Deeper Message",
    })
})

engine.Run()
```

## Example
For more information please see [Example](example).
