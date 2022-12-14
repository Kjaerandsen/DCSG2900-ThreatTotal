import './output.css';
import './App.css';
import Homepage from './pages/homepage';
import Upload from './pages/upload'
import Investigate from './pages/investigate';
import Result from './pages/result';
import About from './pages/about';
import Logout from './pages/logout'
import Login from './pages/login'
import { BrowserRouter as Router, Routes, Route}
    from 'react-router-dom';
//const domain = process.env.REACT_APP_OAUTH_DOMAIN
//const clientID = process.env.REACT_APP_OAUTH_CLIENT_ID

function App() {
  return (
	<Router>
    <Routes>
		<Route path='/' element={<Homepage />} />
    <Route path='/upload' element={<Upload />} />
    <Route path='/investigate' element={<Investigate />} />
    <Route path='/result' element={<Result />} />
    <Route path='/about' element={<About />} />
    <Route path='/logout' element={<Logout />} />
    <Route path='/login' element={<Login />} />
	  </Routes>
	</Router>
  );
}

export default App;
