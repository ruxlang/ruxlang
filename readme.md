## Rux (ruse) designed for simplicity

### A neat new language with a strange name
This language was created from my personal passion for computer
programming and I really liked some of the directions that languages like
[RUST](https://www.rust-lang.org/) and [Go](https://golang.org/) were
taking, however I found many parts of these languages
steeped in antiquated "C" language contraints using naming and namespace
separators like `println` or `io::stdin`.  I get it, they want the basic
language semantics to be familiar to "C" developers.  Personally, I think
we need to dispense with the old baggage already, so I decided to create a
more human friendly language with a unique name.  The main goal was to
create a language a child could pick up without needing to study a dozen
computer science textbooks.  Since the premise of this language is on
interface abstractions like [Go](https://golang.org/) and
[RUST](https://www.rust-lang.org/), I decided to
call it **Rux**, pronounced *ruse* with a soft *x* as in xylophone.
Implementing an interface is like creating a ruse since any component could
imitate that interface.  Thus the name was born.

## How this language is different
OK.  So, we wanted to split from the pack, but why and how much?

Well, the language in [RUST](https://www.rust-lang.org/) is basically C++
if it was rewritten today with immutable by default, mixins, generics,
closures, and removed function definition cruft.  The basic "C" language
constructs are nearly identical like: semicolon line endings, macros,
double colon namespace separators (even typing that hurt my head), etc.

We need these in a modern language, really?

I decided that Go was the closest syntax to my goal for a new language that dispenses
with the old baggage, but the naming (println), pedantic K&R brace handling,
and lack of generics was still an issue for me, not to mention embedding package
management into the source files.  I'm sorry, but tracking where the packages
come from and which version to use isn't the job of source files.  In this
respect I think RUST has it right by handing the responsibility to a package
management tool.  The final straw was no binary library support.
RUST supports binary shared libraries, and any serious language intent on
supporting the development of operating system code must support this.
This leads me to conclude that the Go developers have no intention of growing
Go beyond being a cool monolithic web application development program.

So where do we "Go" from here?

OK, that was a bad pun.  Seriously though.

 * like "C" the language should not be pedantic about brace style.  It is a
   developer codebase preference, not a point of contention that any "C" style
   language should alienate people over.
   If a codebase decides on a style, it should be possible to enforce, but not a
   language requirement.

 * like RUST packaging binary libraries should be possible

 * like Go it should apply all interface definitions at compile time without requiring
   explicit declarations of those implementations
 
 * it should not require any definitions or imports to write a basic program it should
   be as simple as: `func main(){ Print("Hello World") }`