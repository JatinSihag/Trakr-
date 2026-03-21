import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";

interface DashboardData {
    date: string;
    daily_target: number;
    calories_consumed: number;
    calories_burned: number;
    calories_remaining: number;
    status: string;
}

export default function Dashboard() {
    const [data, setData] = useState<DashboardData | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const fetchDashboard = async () => {
            try {
                // 2. THE TOKEN BYPASS: Grab the token manually
                const token = localStorage.getItem('trakr_token');
                console.log("SENDING TOKEN:", token);

                // Force it into the request so Go accepts it
                const response = await api.get('/dashboard', {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });

                setData(response.data);
                setLoading(false);
            } catch (error: any) {
                if (error.response?.status === 401) {
                    handleLogout();
                } else {
                    setError('Failed to load Dashboard');
                    setLoading(false);
                }
            }
        };

        fetchDashboard();
    }, [navigate]);

    const handleLogout = () => {
        localStorage.removeItem('trakr_token');
        navigate('/login');
    }

    if (loading) return <div style={{ textAlign: 'center', marginTop: '100px', fontSize: '18px' }}>Loading your stats...</div>
    if (error) return <div style={{ textAlign: 'center', marginTop: '100px', color: '#dc3545' }}>{error}</div>

    return (
        <div style={{ maxWidth: '800px', margin: '40px auto', padding: '0 20px', fontFamily: 'system-ui, sans-serif' }}>

            <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '40px', paddingBottom: '20px', borderBottom: '1px solid #eee' }}>
                <div>
                    <h1 style={{ margin: '0', color: '#333' }}>Welcome, Champ! 🚀</h1>
                    <p style={{ margin: '5px 0 0 0', color: '#666' }}>Here is your snapshot for {data?.date}</p>
                </div>
                <button
                    onClick={handleLogout}
                    style={{ padding: '8px 16px', backgroundColor: '#f8f9fa', color: '#dc3545', border: '1px solid #dc3545', borderRadius: '6px', cursor: 'pointer', fontWeight: 'bold' }}
                >
                    Logout
                </button>
            </header>

            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))', gap: '20px', marginBottom: '40px' }}>

                {/* Target / Goal Box */}
                <div style={{ backgroundColor: '#f8f9fa', padding: '20px', borderRadius: '12px', boxShadow: '0 2px 4px rgba(0,0,0,0.05)' }}>
                    <h3 style={{ margin: '0 0 10px 0', color: '#555', fontSize: '14px', textTransform: 'uppercase' }}>Daily Target</h3>
                    <p style={{ margin: '0', fontSize: '24px', fontWeight: 'bold', color: '#222' }}>
                        {Math.round(data?.daily_target || 0)} kcal
                    </p>
                    <p style={{ margin: '5px 0 0 0', color: '#666', fontSize: '14px' }}>
                        Goal: {data?.status.replace('_', ' ')}
                    </p>
                </div>

                {/* Food Logged Box */}
                <div style={{ backgroundColor: '#e3f2fd', padding: '20px', borderRadius: '12px', boxShadow: '0 2px 4px rgba(0,0,0,0.05)' }}>
                    <h3 style={{ margin: '0 0 10px 0', color: '#1565c0', fontSize: '14px', textTransform: 'uppercase' }}>Food Logged</h3>
                    <p style={{ margin: '0', fontSize: '24px', fontWeight: 'bold', color: '#0d47a1' }}>
                        {data?.calories_consumed} kcal
                    </p>
                    <p style={{ margin: '5px 0 0 0', color: '#1565c0', fontSize: '14px' }}>Keep it up!</p>
                </div>

                {/* Workouts Box */}
                <div style={{ backgroundColor: '#e8f5e9', padding: '20px', borderRadius: '12px', boxShadow: '0 2px 4px rgba(0,0,0,0.05)' }}>
                    <h3 style={{ margin: '0 0 10px 0', color: '#2e7d32', fontSize: '14px', textTransform: 'uppercase' }}>Workouts</h3>
                    <p style={{ margin: '0', fontSize: '24px', fontWeight: 'bold', color: '#1b5e20' }}>
                        {data?.calories_burned} kcal
                    </p>
                    <p style={{ margin: '5px 0 0 0', color: '#2e7d32', fontSize: '14px' }}>Burned today.</p>
                </div>

            </div>

            <h3 style={{ marginBottom: '15px', color: '#333' }}>Quick Actions</h3>
            <div style={{ display: 'flex', gap: '15px' }}>
                <button style={{ flex: 1, padding: '15px', backgroundColor: '#007bff', color: 'white', border: 'none', borderRadius: '8px', fontSize: '16px', fontWeight: 'bold', cursor: 'pointer' }}>
                    + Log Meal
                </button>
                <button style={{ flex: 1, padding: '15px', backgroundColor: '#28a745', color: 'white', border: 'none', borderRadius: '8px', fontSize: '16px', fontWeight: 'bold', cursor: 'pointer' }}>
                    + Log Workout
                </button>
            </div>
        </div>
    );
}