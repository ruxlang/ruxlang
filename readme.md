## Rux (ruse) designed for simplicity

### A neat new language with a strange name
This language was created partly from my personal passion for computer programming and was
inspired by some of the directions that languages like [RUST](https://www.rust-lang.org/) and
[Go](https://golang.org/) were taking, however I found many parts of these languages still
contained antiquated "C" language constraints using naming and namespace separators like `println`
or `io::stdin`.  I get it, they want the basic language semantics to be familiar to "C"
developers.  Personally, I think it is time to dispense with the old baggage, so I decided
to create a more human friendly language with a unique name.  The main goal was to create a
language anyone could pick up without needing to study a dozen computer science textbooks.

Since the premise of this language is on interface abstractions like [Go](https://golang.org/) and
[RUST](https://www.rust-lang.org/), I decided to call it **Rux**, pronounced *ruse* with a
soft *x* as in xylophone.  Implementing an interface is like creating a ruse since any component
could imitate that interface.  Thus the name was born.  The other premise for this language
is to simplify the complexities with sync/async I/O stream processing. There is little reason
for async I/O to be significantly more complicated than sync I/O and with the increasing number
of processor cores at our disposal the need for performing async processing is only rising
in order to fully utilize them. Rux solves this by first-classing what is called the `pipe`
type for streaming data in some ways similar to a shell pipe which all I/O and iteration
sequences in Rux uses.

## How this language is different
OK.  So, we wanted to split from the pack, but why and how much?

Well, the language in [RUST](https://www.rust-lang.org/) is basically C++ if it was rewritten
today with immutable by default, mixins, generics, closures, and less function definition
cruft.  The basic "C" language constructs are nearly identical like: semicolon line endings, macros,
double colon namespace separators (even typing that hurt my head), etc.

Someone starting with software development will not get all these complexities.  Lets keep the good
parts of immutability, generics, closures, but keep those namespaces simple.  Forget the double
colons and just use simply use the `'.'` for all of them.  This might add minor ambiguity, but this
can easily be resolved in simpler ways like local renaming.

I decided that Go had the closest syntax to my goal for a new language, one that that dispenses
with most of the old baggage so most of the syntax follows similar syntax. Except for the
antiquated naming conventions `Println` for instance, and pedantic K&R brace handling which to
me is really a matter of preference. The module support in Go improved, but still tracking where
the packages come from and which versions to use shouldn't bleed into the source files and changing
this shouldn't break code unless an incompatible change happens. It negatively impacts portability
and creates unnecessary code churn. Rux moves this to a package manager and stores all external
references in the project file instead of the source files. RUST supports binary shared libraries,
and any serious language intent on supporting the development of operating system code must support
this in order to integrate with code you don't own written in other languages, so Rux supports this
as well. Code can be exported as "C" style exports for importing into other languages.

So where do we "Go" from here?

OK, that was a bad pun.  Seriously though.

 * Like "C", the language should not be pedantic about brace style.  It is a developer codebase
   preference, not a point of contention that any "C" style language should alienate people over.
   
 * If a codebase decides on a style, it should be possible to enforce, but not a language
   requirement.  A project wide preference for brace style should be set.  All preferences
   support being overridden where necessary.

 * Like RUST, packaging binary libraries should be possible and specifying library exports to
   "C" based languages made easy.

 * Like Go, it should apply all interface definitions at first cast without requiring explicit
   declarations of those implementations.  This maintains a pseudo interface duck-typing style
   allowing the freedom to apply interfaces against almost any type.  This improves the mocking
   capabilities for unit testing code and allows adding interfaces later without impacting existing
   code.

 * Like Go, it should not need semicolons to end a line or double-colons to reference a namespace,
   this was needed by compilers ages ago, but it isn't necessary anymore.  Instead all namespaces
   are referenced separated by `'.'`.
 
 * Unlike most modern languages these days, it should not require any definitions or
   imports to write a basic program (Auto-imports).

   It should be as simple as:
   ```go
   func main()
   {
       io.PrintLine("Hello World")
   }
   ```

   Or even simpler without `main()` and using the aliased version of `print()`, where the `!`
   indicates to add a newline character to the string.
   ```go
   print("Hello World"!)
   ```

   However, although it is useful for learning and initial experimenting, throwing code in
   the top-level context outside functions is greatly discouraged. It will execute without
   any guaranteed order of execution.  The only guarantee is that it will execute when that
   type is instantiated and in the sequence defined within that file.  The preferred method
   is to use a constructor since all this code is automatically injected into the constructor
   at compile time.  Using a constructor makes it explicitly clear what you are doing.  This
   feature makes the `main()` method unnecessary, except without the order guarantee of using
   `main()`.

 * The `!` newline suffix adds a newline to any string.  Adding a newline is such a common
   pattern that the old `\n` seems more like a hack and doesn't work with string literals
   like this does.  It can be appended to any string type, whereas `\n` is only supported
   in formatted strings (double quotes).
   ```go
   x := `File path C:\testing\this\out, with newline`!
   ```

 * To support scripting and normal C-style commenting `#`, `//`, and `/* */` are all
   supported code commenting styles.  This makes it possible to insert the shebang at the
   top of a source file to be executed by a shell.
   ```go
   #!/bin/rxsh
   // C-style single line
   /*
     C-style multi-line
   */
   print("Hello World"!)
   ```
   This code can be compiled, cached, and executed by the shell.

 * Due to the ubiquity with which the `\` escape character is used in strings across
   languages, it is also supported here with most of the C-style escapes and aliases like
   `\n` and `\r`.  One exception is the addition of the `\l` escape representing linefeed
   (`\x0A`).  This is because `\n` compiles to an architecture-specific definition of a
   new line.  On Windows, `\n` is equivalent to `\r\l` and on Linux `\l`.  This
   automatically preserves line endings so less translation is required when writing
   or reading files and console I/O on different architectures.

 * Like RUST, linking to "C" libraries should be easy, but not require pulling in the whole
   gcc/msvc toolkit.  Keep it simple, build it fast.  Linking also solves some problematic
   licensing issues for libraries with LGPL licenses where Go would build the source code
   directly into the binary, invalidating the linking exception.

 * Automatic "C" header file handling should be available so library exports can be
   imported automatically and map against builtin types wherever possible. Also,
   data marshalling should be automatic to remove most of the guesswork.

 * Like C# 6.0 the `string` builtin type should be a first class citizen of the language
   with formatting built into the design.  Use `fmt.Sprintf()` to format strings?
   Think again!  Format it within the quotes automatically inside braces (`{}`):
   ```go
   y := 32
   x := "y = {y}"
   ```
   
   or specify the format after the colon (`:`) for hexadecimal output:
   ```go
   y := 32
   x := "{y} in hexadecimal is {y:X}"
   ```
   
   and most single line expressions are supported:
   ```go
   y := 3
   x := "{y} + 5 = {y + 5}"
   ```

   All of these are built and verified at compile time, injecting the constant values
   and automatically translated to the string formatting call where required.

 * Like Go all strings are in unicode format to permit use of the extended character
   codes and unicode combining characters while still supporting the basic ASCII text.
   Translation to other code points is available via a library.  Some automatic
   conversions occur when linked against libraries using different character sets.

 * Unlike most languages, all initialized variables are by default immutable because
   this is a safer design pattern, so this is not valid code:
   ```go
   y := 3
   y = 5
   ```

   Or use the `def` keyword which is equivalent to `:=` in this use case
   ```go
   def y = 3
   ```

   You can also use `def` to define the name and type, but provide the value later.
   The value can only be assigned once.
   ```go
   def y int
   y = 5
   ```

 * It is possible to make variables mutable by using the `var` keyword instead, however
   using immutable types is the preferred pattern wherever possible:
   ```go
   var y = 3
   y = 5
   ```

 * Anything in your local project namespace overrides any imported namespace, there is no
   exception.  If you want to import something that collides, alias the imported namespace.
   This is only necessary if some names within your namespace are the same. Generally, it is
   bad practice to use identical names to another package unless they are intended to be used
   separately.  This doesn't prevent you from using the same namespace if you want to add
   features to an existing namespace.

 * Anything under a top-level `internal` path would only be available to the current project,
   but anything in a nested `internal` path would only be visible to its parents up to the top level,
   but not including other top level paths, nor their child namespaces.  This enforces some
   cleanliness in code organization.

 * Keywords or tags in the path can be used to target some specific code to an architecture or
   conditionally build for debug or testing purposes.
   Custom tags can also be added like `integration` to limit building for production use.
   `test` is a keyword tag which is automatically set when running tests.

 * Like C#, Java, and Kotlin, package names and namespaces should be named in such a
   way so that they can be replaced with identical versions without rewriting the code.
   So an import definition would not use the path, but a dot separated namespace
   like this:
   ```go
   import lexer.compiler
   ```

 * Paths in your codebase or external projects are automatically treated like
   namespaces unless a package name overrides this.  Also, anything under `internal` or
   a keyword path like `windows`, `linux`, `amd64` or `arm64` will be excluded from the
   namespace unless the package name overrides this.  These special namespaces are compile
   keywords which limit building these paths unless a matching tag or build architecture is set.
   The file path `internal/lexer/compiler` would be imported like this:
   ```go
   import lexer.compiler
   ```
   or the path `os/windows/path` is imported as:
   ```go
   import os.path
   ```
   or if the `integration` tag is set, the path `internal/app/integration/web` becomes:
   ```go
   import app.web
   ```
   Notice the fact that although it is an internal component the `internal` path does not affect
   the namespace.  This is important for code where you may want to make an internal component
   public eventually, but only if it can be reused.  It would be a pain to rewrite imports
   in working code only to make a package public to new code.  It also means that the opposite
   is also true, if you make something internal later it is not an internal breaking change,
   just an external one.

 * A path excluded by a tag should usually not be repeated in any given namespace path, since only the
   first occurrence will be stripped from the namespace, any subsequent repetitions will become part
   of the namespace.  Although this could be intentional if you want the name to be part of the
   namespace.
   An example could be the path `os/windows/windows/registry` which becomes:
   ```go
   import os.windows.registry
   ```
   In this example we are referring to an OS specific feature, but it will only compile for that OS
   target as well.  This form automatically encourages architecture specific code to be segmented.
   The alternative would be to use the path `os/windows/registry` and set the package name inside
   `registry`, but this adds work and limits the namespace change to `registry` only, so it isn't
   automatic nor the preferred pattern unless `registry` is the only OS specific feature in the
   `windows` namespace.
   You could also set the package name inside `windows`, however you may want to create common
   `os.*` package features within `os/windows/*` (like `os.env`) with OS dependent implementations.
   Having `registry` with a namespace override could get confusing in this case.  Also, forcing the
   `windows` namespace name would make it `os.windows.env`, which is undesirable.  So, it is only
   recommended when you want to change the namespace to inject into the namespace of another package.

 * The package `rux.toml` configuration would import any path with its default namespace
   which can optionally be aliased.
   ```ini
   base=github.com/ruxlang/ruxlang # defines the root path and namespace of the project

   [build]
   _=core@v0.1 # Imports the rux core 0.1 namespace and removes 'core' from the namespace so core.io.file
               # becomes just io.file, the FQDN part of the path is optional for rux core packages
               # Anything originally in the 'core' namespace becomes a built-in like "print"
               # If the project file or config section is missing, this is the default
   lexer=github.com/antlr/antlr4.lexer@v4.7.2 # Imports only antlr4.lexer from the antlr4 namespace
                                              # This will also alias the namespace antlr4.lexer to lexer
   [test]
   test=core.test@v0.1 # Imports test core package namespace as a test dependency
                       # If the project file or config section is missing, this is the default
   ```
   Testing dependencies can be added as well and will be imported only when building tests.
   Segregating these dependencies improves build times.  You can basically add any dependency
   sections for each custom keyword or tag.  Any dependencies not mentioned in the project
   file that those in the project file depend upon will be imported at build time so in most
   cases only direct dependencies should need to be mentioned this keeps the project file
   small and managable.
