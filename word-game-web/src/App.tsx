import { useEffect, useState } from "react";
import "./App.css";

type Color = "green" | "yellow" | "gray";

interface LetterResult {
  letter: string;
  color: Color;
}

interface GuessResponse {
  board: LetterResult[][];
  isOver: boolean;
  won: boolean;
}

interface StartResponse {
  id: string;
  expiresAt: string;
}

const keyboardLayout = [
  ["q", "w", "e", "r", "t", "y", "u", "i", "o", "p"],
  ["a", "s", "d", "f", "g", "h", "j", "k", "l"],
  ["Enter", "z", "x", "c", "v", "b", "n", "m", "Backspace"],
];

function App() {
  const [hasStarted, setHasStarted] = useState(false);
  const [gameId, setGameId] = useState<string | null>(null);
  const [board, setBoard] = useState<LetterResult[][]>([]);
  const [guess, setGuess] = useState("");
  const [isOver, setIsOver] = useState(false);
  const [expiresAt, setExpiresAt] = useState<number | null>(null);
  const [timeLeft, setTimeLeft] = useState("");
  const [error, setError] = useState<string | null>(null);

  const startGame = async () => {
    setError(null); // clear any previous error

    try {
      const res = await fetch("/api/start", { method: "POST" });

      if (!res.ok) {
        const text = await res.text();
        setError(text || "An unknown error occurred.");
        return;
      }

      const data: StartResponse = await res.json();
      setGameId(data.id);
      setBoard([]);
      setGuess("");
      setIsOver(false);
      setHasStarted(true);
      setExpiresAt(new Date(data.expiresAt).getTime());
    } catch (err) {
      console.error(err);
      setError("Failed to contact server. Please try again.");
    }
  };

  useEffect(() => {
    if (!expiresAt) return;
    const interval = setInterval(() => {
      const now = Date.now();
      const diff = expiresAt - now;

      if (diff <= 0) {
        setTimeLeft("00:00");
        clearInterval(interval);
        setIsOver(true);
        return;
      }

      const minutes = Math.floor(diff / 60000);
      const seconds = Math.floor((diff % 60000) / 1000);
      setTimeLeft(
        `${String(minutes).padStart(2, "0")}:${String(seconds).padStart(
          2,
          "0"
        )}`
      );
    }, 1000);

    return () => clearInterval(interval);
  }, [expiresAt]);

  const handleKey = (key: string) => {
    if (isOver) return;
    if (key === "Enter") {
      submitGuess();
    } else if (key === "Backspace") {
      setGuess((prev) => prev.slice(0, -1));
    } else if (guess.length < 5 && /^[a-z]$/.test(key)) {
      setGuess((prev) => prev + key);
    }
  };

  const submitGuess = async () => {
    if (!gameId || guess.length !== 5) return;

    const res = await fetch("/api/guess", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id: gameId, guess }),
    });

    const data: GuessResponse = await res.json();
    setBoard(data.board);
    setGuess("");
    setIsOver(data.isOver);

    if (data.isOver) {
      setTimeout(() => {
        alert(data.won ? "🎉 You won!" : "❌ You lost!");
      }, 50);
    }
  };

  const renderTile = (tile: Partial<LetterResult> = {}, key: number) => (
    <div className={`tile ${tile.color ?? "input"}`} key={key}>
      {tile.letter?.toUpperCase?.() ?? ""}
    </div>
  );

  if (!hasStarted) {
    return (
      <div className="app landing">
        <h1>Welcome to Word Guessing Game</h1>
        <button className="start-btn" onClick={startGame}>
          ▶ Start Game
        </button>
        {error && <p className="error-msg">{error}</p>}
      </div>
    );
  }

  return (
    <div className="app">
      <h1>Word Guessing Game</h1>

      {expiresAt && <p className="timer">⏳ Time left: {timeLeft}</p>}

      <div className="board">
        {board.map((row, i) => (
          <div className="row" key={i}>
            {row.map((tile, j) => renderTile(tile, j))}
          </div>
        ))}

        {!isOver && (
          <div className="row">
            {Array.from({ length: 5 }).map((_, i) => {
              const letter = guess[i] ?? "";
              return renderTile({ letter }, i);
            })}
          </div>
        )}
      </div>

      {isOver && (
        <button className="new-game-btn" onClick={startGame}>
          🔄 New Game
        </button>
      )}

      <div className="keyboard">
        {keyboardLayout.map((row, i) => (
          <div className="keyboard-row" key={i}>
            {row.map((key) => (
              <button
                className={`key ${
                  key === "Enter" || key === "Backspace" ? "special" : ""
                }`}
                key={key}
                onClick={() => handleKey(key)}
              >
                {key === "Backspace" ? "←" : key.toUpperCase()}
              </button>
            ))}
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
