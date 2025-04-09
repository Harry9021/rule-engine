import { useState } from "react";
import axios from "axios";

const RuleForm = () => {
  const [id, setId] = useState("");
  const [condition, setCondition] = useState("");
  const [action, setAction] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await axios.post("http://localhost:8080/rules", {
        id,
        condition,
        action,
      });
      setMessage("✅ Rule created successfully!");
      setId("");
      setCondition("");
      setAction("");
      setTimeout(() => {
        window.location.reload();
      }, 500);
    } catch (err) {
      console.error(err);
      setMessage("❌ Failed to create rule.");
    }
  };

  return (
    <div className="bg-white p-6 rounded-xl shadow-md mb-6">
      <h2 className="text-xl font-semibold mb-4">➕ Add New Rule</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="text"
          placeholder="Rule ID"
          value={id}
          onChange={(e) => setId(e.target.value)}
          className="w-full p-2 border rounded"
          required
        />
        <input
          type="text"
          placeholder="Condition (e.g., temp > 40)"
          value={condition}
          onChange={(e) => setCondition(e.target.value)}
          className="w-full p-2 border rounded"
          required
        />
        <input
          type="text"
          placeholder="Action (e.g., alert('High Temp'))"
          value={action}
          onChange={(e) => setAction(e.target.value)}
          className="w-full p-2 border rounded"
          required
        />
        <button
          type="submit"
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 cursor-pointer"
        >
          Create Rule
        </button>
        {message && <p className="text-sm text-green-600">{message}</p>}
      </form>
    </div>
  );
};

export default RuleForm;
