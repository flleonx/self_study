## Packing

The modules map is the final output of the **_dependency resolution phase_**.
In the packing phase, the module bundler takes the modules map and converts it
into an _executable bundle_: a single JavaScript file that contains all the
business logic of the original application.

caculator.js example after dependency graph resolution:
```javascript
    (module, require) => {
        const { parser } = require('parser.js');
        const { resolver } = require('resolver.js');

        module.exports.calculator = function (expr) {
            return resolver(parser(expr))
        }
    }
```

```javascript
    ((modulesMap) => {
        const require = (name) => {
            const module = { exports: {} }
            modulesMap[name](module, require)
            return module.exports
        }

        require('app.js')
    })(
        {
            'app.js': (module, require) => {/* ... */},
            'calculator.js': (module, require) => {/* ... */},
            'display.js': (module, require) => {/* ... */},
            'parser.js': (module, require) => {/* ... */},
            'resolver.js': (module, require) => {/* ... */}, 
        }
    )
```


NOTE: Finally, you get a file with an object which contains all the dependecies
resolved in a proper way.
