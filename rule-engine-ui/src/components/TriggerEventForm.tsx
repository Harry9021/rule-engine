import { useState } from "react";
import axios from "axios";

const TriggerEventForm = () => {
  const [eventData, setEventData] = useState<string>('{\n  "temp": 45\n}');
  const [message, setMessage] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const parsedData = JSON.parse(eventData);
      await axios.post("http://localhost:8080/event", parsedData);
      setMessage("✅ Event triggered successfully!");
    } catch (err) {
      console.error(err);
      setMessage("❌ Failed to send event. Check your JSON.");
    }
  };

  return (
    <div className="bg-white p-6 rounded-xl shadow-md">
      <h2 className="text-xl font-semibold mb-4">🚀 Trigger Event</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <textarea
          rows={6}
          value={eventData}
          onChange={(e) => setEventData(e.target.value)}
          className="w-full p-2 border rounded font-mono"
        />
        <button
          type="submit"
          className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
        >
          Send Event
        </button>
        {message && <p className="text-sm mt-2">{message}</p>}
      </form>
    </div>
  );
};

export default TriggerEventForm;
