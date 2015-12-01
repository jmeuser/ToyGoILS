# ToyGoILS
A toy ILS written in Go.

$ go run main.go

then navigate to http://localhost:8080/view?c=UniCat and you should see:



<h1>UniCat</h1>


<div>
<table>
<tr>
 <th scope="row">Title</th>
 <td>Authority and the Individual</td>
</tr>
<tr>
 <th scope="row">ISBN</th>
 <td>9781134812271</td>
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

</div>

<div>
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

</div>



To Do:
* Introduce edit book
* Introduce view catalogue
* Introduce search for book
* Introduce search for book in catalogue
* Introduce homepage
* Eliminate possible race conditions using "synch" package and synch.Mutex or related tool.
* More to Come
