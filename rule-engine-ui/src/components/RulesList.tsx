import { useEffect, useState } from "react";
import axios from "axios";

type Rule = {
  id: string;
  condition: string;
  action: string;
};

const RulesList = () => {
  const [rules, setRules] = useState<Rule[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchRules = async () => {
    try {
      const res = await axios.get("http://localhost:8080/rules");
      setRules(res.data);
    } catch (err) {
      console.error("Failed to fetch rules:", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchRules();
  }, []);

  return (
    <div className="bg-white p-6 rounded-xl shadow-md mb-6">
      <h2 className="text-xl font-semibold mb-4">📋 Existing Rules</h2>
      {loading ? (
        <p>Loading...</p>
      ) : rules.length === 0 ? (
        <p>No rules found.</p>
      ) : (
        <table className="w-full table-auto border">
          <thead>
            <tr className="bg-gray-100 text-left">
              <th className="p-2 border">ID</th>
              <th className="p-2 border">Condition</th>
              <th className="p-2 border">Action</th>
            </tr>
          </thead>
          <tbody>
            {rules.map((rule) => (
              <tr key={rule.id}>
                <td className="p-2 border">{rule.id}</td>
                <td className="p-2 border">{rule.condition}</td>
                <td className="p-2 border">{rule.action}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default RulesList;
