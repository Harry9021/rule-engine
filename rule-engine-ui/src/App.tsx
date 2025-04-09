import RuleForm from "./components/RuleForm";
import RulesList from "./components/RulesList";
import TriggerEventForm from "./components/TriggerEventForm";
// import SystemMonitor from "./components/SystemMonitor";

const App = () => {
  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <div className="max-w-3xl mx-auto">
        <h1 className="text-3xl font-bold mb-6 text-center">⚙️ Rule Engine UI</h1>
        {/* <SystemMonitor /> */}
        
        <RuleForm />
        <RulesList />
        <TriggerEventForm />
      </div>
    </div>
  );
};

export default App;
