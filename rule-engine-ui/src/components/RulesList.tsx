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
  const [editingRule, setEditingRule] = useState<Rule | null>(null);
  const [editCondition, setEditCondition] = useState("");
  const [editAction, setEditAction] = useState("");

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

  const handleEdit = (rule: Rule) => {
    setEditingRule(rule);
    setEditCondition(rule.condition);
    setEditAction(rule.action);
  };

  const handleCancelEdit = () => {
    setEditingRule(null);
    setEditCondition("");
    setEditAction("");
  };

  const handleSaveEdit = async () => {
    if (!editingRule) return;

    try {
      await axios.put("http://localhost:8080/rules", {
        id: editingRule.id,
        condition: editCondition,
        action: editAction,
      });
      
      // Update local state
      setRules(rules.map(rule => 
        rule.id === editingRule.id 
          ? { ...rule, condition: editCondition, action: editAction } 
          : rule
      ));
      
      // Reset editing state
      setEditingRule(null);
      setEditCondition("");
      setEditAction("");
    } catch (err) {
      console.error("Failed to update rule:", err);
      alert("Failed to update rule. Please try again.");
    }
  };

  const handleDelete = async (id: string) => {
    if (!confirm("Are you sure you want to delete this rule?")) return;

    try {
      await axios.delete(`http://localhost:8080/rules?id=${id}`);
      
      // Update local state
      setRules(rules.filter(rule => rule.id !== id));
    } catch (err) {
      console.error("Failed to delete rule:", err);
      alert("Failed to delete rule. Please try again.");
    }
  };

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
              <th className="p-2 border">Operations</th>
            </tr>
          </thead>
          <tbody>
            {rules.map((rule) => (
              <tr key={rule.id}>
                <td className="p-2 border">{rule.id}</td>
                <td className="p-2 border">
                  {editingRule?.id === rule.id ? (
                    <input
                      type="text"
                      value={editCondition}
                      onChange={(e) => setEditCondition(e.target.value)}
                      className="w-full p-1 border rounded"
                    />
                  ) : (
                    rule.condition
                  )}
                </td>
                <td className="p-2 border">
                  {editingRule?.id === rule.id ? (
                    <input
                      type="text"
                      value={editAction}
                      onChange={(e) => setEditAction(e.target.value)}
                      className="w-full p-1 border rounded"
                    />
                  ) : (
                    rule.action
                  )}
                </td>
                <td className="p-2 border">
                  {editingRule?.id === rule.id ? (
                    <div className="flex space-x-2">
                      <button
                        onClick={handleSaveEdit}
                        className="bg-green-500 text-white px-2 py-1 rounded hover:bg-green-600"
                      >
                        Save
                      </button>
                      <button
                        onClick={handleCancelEdit}
                        className="bg-gray-500 text-white px-2 py-1 rounded hover:bg-gray-600"
                      >
                        Cancel
                      </button>
                    </div>
                  ) : (
                    <div className="flex space-x-2">
                      <button
                        onClick={() => handleEdit(rule)}
                        className="bg-blue-500 text-white px-2 py-1 rounded hover:bg-blue-600"
                      >
                        Edit
                      </button>
                      <button
                        onClick={() => handleDelete(rule.id)}
                        className="bg-red-500 text-white px-2 py-1 rounded hover:bg-red-600"
                      >
                        Delete
                      </button>
                    </div>
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default RulesList;
