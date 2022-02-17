import logo from './logo.svg';
import './output.css';
import './App.css'
import Navbar from './navbar'
import Indextest from './indextest';
import { BrowserRouter as Router, Routes, Route}
    from 'react-router-dom';

function App() {
  return (
	<Router>
    <Routes>
		<Route exact path='/' exact element={<Indextest />} />
	</Routes>
	</Router>
  );
}

export default App;
