import './output.css';
import './App.css';
import Upload from './pages/upload'
import Indextest from './pages/indextest';
import Investigate from './pages/investigate';
import Result from './pages/result';
import About from './pages/about';
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
