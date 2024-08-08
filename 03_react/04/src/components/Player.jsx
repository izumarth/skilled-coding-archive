import { useState, useRef } from 'react';

export default function Player() {
  const playerName = useRef();
  const [enterdPlalyerName, setEnterdPlalyerName] = useState(null);

  function handleClick() {
    setEnterdPlalyerName(playerName.current.value);
    playerName.current.value == '';
  }

  return (
    <section id="player">
      <h2>Welcome {enterdPlalyerName ?? 'unknow entity'}</h2>
      <p>
        <input ref={playerName} type="text"/>
        <button onClick={handleClick}>Set Name</button>
      </p>
    </section>
  );
}
