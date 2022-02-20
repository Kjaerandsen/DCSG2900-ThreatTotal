import './output.css';
import './App.css';
import Upload from './upload'
import Indextest from './indextest';
import { BrowserRouter as Router, Routes, Route}
    from 'react-router-dom';

function App() {
  return (
	<Router>
    <Routes>
		<Route path='/' element={<Indextest />} />
    <Route path='/upload' element={<Upload />} />
	  </Routes>
	</Router>
  );
}

export default App;
