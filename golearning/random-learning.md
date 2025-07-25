## Interfaces In GO
- Defining declaration without definition we define interfaces
- Methods declaraed in interfaces are contracts that srtuct can implement
- Grouping of methods in Interfaces can provide group set of functionality
- In Go as long as methods are implemented in structs. Struct indirectly implements that particular interface


# Go Gotchas Cheat Sheet

## ✅ Export Rules (Structs, Methods, Fields)
- **Capitalized identifiers** → **Exported (public)**, accessible from other packages.
- **Lowercase identifiers** → **Unexported (private)**, accessible only within the same package.
- Both **struct name** and **field names** must be capitalized to use them from other packages.
- A **lowercase struct name** is completely inaccessible outside its package, even if its fields are capitalized.

---

## ✅ Interfaces & Methods
- A struct automatically implements an interface if it defines all methods with **matching names (case-sensitive)** and **signatures**.
- Go does **not** support method overloading (no two methods with the same name, even if signatures differ).
- **Two interfaces with the same name** cannot exist in the same package.
- Interfaces can be **composed** (embedded):

  ```go
  type Reader interface { Read() }
  type Writer interface { Write() }
  type ReadWriter interface {
      Reader
      Writer
  }
