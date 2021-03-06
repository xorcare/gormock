# [GORMock] / COMP

The package contains structures with interfaces compatible with [GORM]
structures.

## Checklist

 * Implement interfaces
    * [x] gorm.DB ≈ comp.DB 
    * [x] gorm.Scope ≈ comp.Scope 
    * [x] gorm.Association ≈ comp.Association
    * [x] gorm.Callback
    * [x] gorm.CallbackProcessor ≈ comp.CallbackProcessor
    * [ ] gorm.Dialect
    * [ ] gorm.Field
    * [ ] gorm.JoinTableHandlerInterface
    * [x] gorm.Search ≈ comp.Search
    * [ ] gorm.SQLCommon

## License

© Vasiliy Vasilyuk, 2018-2019

Released under the [BSD 3-Clause License][LICENSE].

[GORMock]: https://git.io/fhHpT 'The fantastic mock for the fantastic GORM
library, aims to be developer friendly.'
[GORM]: https://git.io/fhHbK 'GORM The fantastic ORM library for Golang, aims
to be developer friendly.'
[LICENSE]: https://git.io/fhHbM 'BSD 3-Clause "New" or "Revised" License'
