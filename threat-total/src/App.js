import './output.css';
import './App.css';
import Upload from './upload'
import Indextest from './indextest';
import Investigate from './investigate';
import { BrowserRouter as Router, Routes, Route}
    from 'react-router-dom';

function App() {
  return (
	<Router>
    <Routes>
		<Route path='/' element={<Indextest />} />
    <Route path='/upload' element={<Upload />} />
    <Route path='/investigate' element={<Investigate />} />
	  </Routes>
	</Router>
  );
}

export default App;
