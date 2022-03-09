import './output.css';
import './App.css';
import Upload from './upload'
import Indextest from './indextest';
import Investigate from './investigate';
import Result from './result';
import About from './about';
import { BrowserRouter as Router, Routes, Route}
    from 'react-router-dom';

function App() {
  return (
	<Router>
    <Routes>
		<Route path='/' element={<Indextest />} />
    <Route path='/upload' element={<Upload />} />
    <Route path='/investigate' element={<Investigate />} />
    <Route path='/result' element={<Result />} />
    <Route path='/about' element={<About />} />
	  </Routes>
	</Router>
  );
}

export default App;
