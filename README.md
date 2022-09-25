# Yak shop

Empowering the local yak farmer.

This service uses the following endpoints:

---

### Stock

`GET` `http://localhost:8080/yak-shop/stock/14`

Gets information about the store's stock, for the given day.

---

### Herd

`GET` `http://localhost:8080/yak-shop/herd/13`

Gets information about the state of the herd, for the given day.

---


### Order

`POST` `http://localhost:8080/yak-shop/order/14`

```json
{
	"customer": "Sample customer",
	"order": {
		"milk": 1100,
		"skins": 3
	}
}
```

Posts an order to the shop. Orders can be partially fulfilled. If no goods are available in enough quantity, the order fails.

---

#### Technical info

This is a Go API application that imports data from a XML file and stores this data in memory, in an immutable fashion. For every successful request, the data is modified according to the specified `day` attribute (present in all possible API calls) and returns the desired information about the current state of the herd, specific to this day.

To specify a different XML file for importing, set the `-file` flag when running this application. Check the XML syntax below.

A `herd-sample.xml` file is present to show how data can be provided. The data has the following format:

```xml
<herd>
    <labyak name="Betty-1" age="4" sex="f"/>
    <labyak name="Betty-2" age="8" sex="f"/>
    <labyak name="Betty-3" age="9.5" sex="f"/>
</herd>
```