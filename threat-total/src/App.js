import './output.css';
import './App.css';
import Upload from './pages/upload'
import Indextest from './pages/indextest';
import Investigate from './pages/investigate';
import Result from './pages/result';
import About from './pages/about';
import { BrowserRouter as Router, Routes, Route}
    from 'react-router-dom';
import "react-loader-spinner/dist/loader/css/react-spinner-loader.css";
//const domain = process.env.REACT_APP_OAUTH_DOMAIN
//const clientID = process.env.REACT_APP_OAUTH_CLIENT_ID

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
