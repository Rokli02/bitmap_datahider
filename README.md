# BMP kiterjesztésű képeken történő adatrejtés

## Leírás

A feladat egy olyan szoftver létrehozása, amely képes szöveget `.bmp` kiterjesztésű fájlokba beágyazni úgy, hogy a lehető legkevésbé rontsa a kép minőségét. A steganográfiában használatos `LSB` - __*Least Significant Bit*__ eljárást használja a program.<br/>
Az alkalmazás kiszámítja, hogy milyen nagyságú byte tömbnek felel meg a megadott szöveg, majd az ebben szereplő biteket addig pakolja a kép *LSB*-jeibe, míg vagy a kép vagy a szöveg bytejai elfogynak. Lehetőség van megadni, hogy byteonként mennyi bitet használjon fel (`BITS_USED_FOR_HIDE`), így minimálisan lehet szabályozni a kép minőségének romlását.
Amennyiben a képben nem történt szöveg eljrejtés, akkor azt a program észreveszi és közli a felhasználóval. Ez azért lehetséges, mert a *.bmp* fájlokban jelen van egy üresen hagyott 4. byte az RGB után. Ebben is történhetne adatrejtés és ebben az esetben gyakorlatilag kép romlás sem történne, de jelen esetben csak metaadat tárolás történik itt. <br/>
A kép amiben az adat el lett rejtve az `assets/sus/` mappába kerül behelyezésre.

## Követelmények

Képesnek kell lennie az alábbi pontokra:
  - Szöveg beágyazása egy BMP fájlba.
  - Szöveg kiolvasása egy BMP fájlból.
  - A beágyazott adat méretének és elérhetőségének ellenőrzése.

## Használat

### Build

Fontos, hogy az alkalmazás futtatásához szükség van __GoLang__ compiler meglétére a számítógépen.
Lehetőség van a kód `.exe` fájlra történő lefordítására:
```
go build -o bmphider.exe .
```

ezt követően az *exe* fájlt kell futtatni, vagy pedig minden futtatáskor frissen lefordítani a kódot:
```
go run .
```

### Futtatás

A build módszertől függetlenül az alkalmazást argumentumok megadásával tudjuk kezelni.
Két fő parancs van: `hide` valamint `reveal`.<br/>
A __*hide*__ parancsnak 2 argumentumot kell megadni az alábbi sorrendben
  - <u>kép elérési útvonala</u>: Az `assets/` mappához képest relatív kell megadni.
  - <u>elrejtendő szöveg</u>: Az a szöveg, amit el szeretnénk rejteni a megadott képben.

```
bmphider.exe hide "mario.bmp" "Elrejtendő adat"
```

A __*reveal*__ parancsnak 1 argumentumot lehet megadni:
  - <u>szöveg felfedése a képből</u>: Az `assets/` mappához képest relatív kell megadni.

```
bmphider.exe reveal "sus/mario.bmp"
```