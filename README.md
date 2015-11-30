# ToyGoILS
A toy ILS written in Go.

$ go run main.go

then navigate to http//localhost:8080/ and you should see:


<table>
<tr>
 <th scope="row">Title</th>
 <td>The Principles of Mathematics</td>
</tr>
<tr>
 <th scope="row">ISBN</th>
 <td>9780203864760</td>
</tr>
<tr>
 <th scope="row">Library</th>
 <td>Pembrook Public Library</td>
</tr>
<tr>
 <th scope="row">Requests</th>
 <td>0</td>
</tr>
</table>

To Do:
* Introduce edit book
* Introduce view catalogue
* Introduce search for book
* Introduce search for book in catalogue
* Introduce homepage
* Eliminate possible race conditions using "synch" package and synch.Mutex or related tool.
* More to Come
