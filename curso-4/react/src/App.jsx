import logo from './logo.svg';
import logogo from './logo-go.png'
import './App.css';
import Personalidades from './components/Personalidades';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logogo} className="App-logo" alt="logo" />
        <h1>Personalidades</h1>
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <Personalidades />
    </div>
  );
}

export default App;
