# ToyGoILS
A toy ILS written in Go.

$ go run main.go


then navigate to http://localhost:8080/find

and put "Authority and the Individual" (without quotes) into the Title text input and click the button labeled "Find".
You should see this:

- - -
<h1>Results</h1>
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
- - -

Go back to http://localhost:8080/find and in the ISBN text input put "9780203864760" (without quotes) and you should see this:

- - -
<h1>Results</h1>

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
- - -

To Do:
* Introduce edit book
* Introduce request book
* Introduce "make" book
* Introduce homepage
* Eliminate possible race conditions using "synch" package and synch.Mutex or related tool so that it would work in "real time" without odd behavior.
* More to Come
