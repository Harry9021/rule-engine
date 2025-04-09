import { useEffect, useState } from "react";
import axios from "axios";

type SystemStats = {
  cpuUsage: number;
  memoryUsage: number;
  timestamp: number;
};

type AlertThreshold = {
  cpuThreshold: number;
  memoryThreshold: number;
};

const SystemMonitor = () => {
  const [stats, setStats] = useState<SystemStats | null>(null);
  const [thresholds, setThresholds] = useState<AlertThreshold>({
    cpuThreshold: 80,
    memoryThreshold: 80,
  });
  const [isEditing, setIsEditing] = useState(false);
  const [tempThresholds, setTempThresholds] = useState<AlertThreshold>({
    cpuThreshold: 80,
    memoryThreshold: 80,
  });

  // Fetch current system stats
  const fetchStats = async () => {
    try {
      const res = await axios.get("http://localhost:8080/monitoring/stats");
      setStats(res.data);
    } catch (err) {
      console.error("Failed to fetch system stats:", err);
    }
  };

  // Fetch current alert thresholds
  const fetchThresholds = async () => {
    try {
      const res = await axios.get(
        "http://localhost:8080/monitoring/thresholds"
      );
      setThresholds(res.data);
      setTempThresholds(res.data);
    } catch (err) {
      console.error("Failed to fetch alert thresholds:", err);
    }
  };

  // Update alert thresholds
  const updateThresholds = async () => {
    try {
      await axios.post(
        "http://localhost:8080/monitoring/thresholds",
        tempThresholds
      );
      setThresholds(tempThresholds);
      setIsEditing(false);
    } catch (err) {
      console.error("Failed to update alert thresholds:", err);
      alert("Failed to update thresholds. Please try again.");
    }
  };

  // Start editing thresholds
  const handleEdit = () => {
    setTempThresholds({ ...thresholds });
    setIsEditing(true);
  };

  // Cancel editing
  const handleCancel = () => {
    setTempThresholds({ ...thresholds });
    setIsEditing(false);
  };

  // Format timestamp to readable date/time
  const formatTimestamp = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleTimeString();
  };

  // Determine color based on usage vs threshold
  const getUsageColor = (usage: number, threshold: number) => {
    if (usage >= threshold) return "text-red-600";
    if (usage >= threshold * 0.8) return "text-yellow-600";
    return "text-green-600";
  };

  useEffect(() => {
    // Initial fetch
    fetchStats();
    fetchThresholds();

    // Set up polling for stats
    const interval = setInterval(fetchStats, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="bg-white p-6 rounded-xl shadow-md mb-6">
      <h2 className="text-xl font-semibold mb-4">🖥️ System Monitor</h2>

      {stats ? (
        <div className="mb-6">
          <div className="grid grid-cols-2 gap-4">
            <div className="border rounded-lg p-4">
              <h3 className="text-lg font-medium mb-2">CPU Usage</h3>
              <div className="flex items-center">
                <div className="w-full bg-gray-200 rounded-full h-4 mr-4">
                  <div
                    className={`h-4 rounded-full ${getUsageColor(
                      stats.cpuUsage,
                      thresholds.cpuThreshold
                    )}`}
                    style={{
                      width: `${Math.min(stats.cpuUsage, 100)}%`,
                      backgroundColor:
                        stats.cpuUsage >= thresholds.cpuThreshold
                          ? "#ef4444"
                          : "#10b981",
                    }}
                  ></div>
                </div>
                <span
                  className={`font-bold ${getUsageColor(
                    stats.cpuUsage,
                    thresholds.cpuThreshold
                  )}`}
                >
                  {stats.cpuUsage.toFixed(1)}%
                </span>
              </div>
              <p className="text-sm text-gray-500 mt-1">
                Threshold: {thresholds.cpuThreshold}%
              </p>
            </div>

            <div className="border rounded-lg p-4">
              <h3 className="text-lg font-medium mb-2">Memory Usage</h3>
              <div className="flex items-center">
                <div className="w-full bg-gray-200 rounded-full h-4 mr-4">
                  <div
                    className={`h-4 rounded-full ${getUsageColor(
                      stats.memoryUsage,
                      thresholds.memoryThreshold
                    )}`}
                    style={{
                      width: `${Math.min(stats.memoryUsage, 100)}%`,
                      backgroundColor:
                        stats.memoryUsage >= thresholds.memoryThreshold
                          ? "#ef4444"
                          : "#10b981",
                    }}
                  ></div>
                </div>
                <span
                  className={`font-bold ${getUsageColor(
                    stats.memoryUsage,
                    thresholds.memoryThreshold
                  )}`}
                >
                  {stats.memoryUsage.toFixed(1)}%
                </span>
              </div>
              <p className="text-sm text-gray-500 mt-1">
                Threshold: {thresholds.memoryThreshold}%
              </p>
            </div>
          </div>

          <p className="text-sm text-gray-500 mt-2">
            Last updated: {formatTimestamp(stats.timestamp)}
          </p>
        </div>
      ) : (
        <p>Loading system stats...</p>
      )}

      <div className="mt-6 border-t pt-4">
        <h3 className="text-lg font-medium mb-3">Alert Thresholds</h3>

        {isEditing ? (
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                CPU Usage Threshold (%)
              </label>
              <input
                type="number"
                min="0"
                max="100"
                value={tempThresholds.cpuThreshold}
                onChange={(e) =>
                  setTempThresholds({
                    ...tempThresholds,
                    cpuThreshold: Number(e.target.value),
                  })
                }
                className="border rounded p-2 w-full"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Memory Usage Threshold (%)
              </label>
              <input
                type="number"
                min="0"
                max="100"
                value={tempThresholds.memoryThreshold}
                onChange={(e) =>
                  setTempThresholds({
                    ...tempThresholds,
                    memoryThreshold: Number(e.target.value),
                  })
                }
                className="border rounded p-2 w-full"
              />
            </div>

            <div className="flex space-x-2">
              <button
                onClick={updateThresholds}
                className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
              >
                Save Thresholds
              </button>
              <button
                onClick={handleCancel}
                className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
              >
                Cancel
              </button>
            </div>
          </div>
        ) : (
          <div>
            <p className="mb-4">
              Set thresholds to receive alerts when system resource usage
              exceeds these values.
            </p>
            <button
              onClick={handleEdit}
              className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
            >
              Edit Thresholds
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default SystemMonitor;
