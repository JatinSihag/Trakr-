import { useState } from "react";
import api from "../services/api";


export default function SignUp() {
    const [name , setName] = useState('');
    const [email, setEmail] = useState('');
    const [password,setPassword] = useState('');
    const [message,setMessage] = useState('');

    const handleSignup= async (e: React.FormEvent)=>{
        e.preventDefault();
        setMessage('');
        try{
            const response = await api.post("/signup",{
                name,
                email,
                password
            })
            setMessage('User Created Successfully!')
            setName('');
            setEmail('');
            setPassword('');
        }catch (error : any){
            setMessage(error.response?.data?.error || 'Something went wrong')
        }
    };

    return (
        <div style={{ maxWidth:"400px",margin:"50px auto",fontFamily:"sans-serif"}}>
            <h2>Join Trakr</h2>
            {message && (
                <div style={{padding:"10px",marginBottom:"15px",backgroundColor:'#f0f0f0',borderRadius:"5px"}}>
                    {message}
                    </div>
            )}
            <form onSubmit={handleSignup} style={{ display:"flex",flexDirection:"column",gap:"15px"}}>
                <input
                type="text"
                placeholder="Enter your name"
                value={name}
                onChange={(e)=>setName(e.target.value)}
                required
                style={{ padding:"10px",fontSize:"16px"}}
                />
                <input
                type="text"
                placeholder="Enter your Email"
                value={email}
                onChange={(e)=>setEmail(e.target.value)}
                required
                style={{ padding: '10px', fontSize: '16px' }}
                />
                <input
                type="password"
                placeholder="Enter your password"
                value={password}
                onChange={(e)=>setPassword(e.target.value)}
                required
                style={{padding: '10px', fontSize: '16px' }}
                />
                <button
                type="submit"
                style={{padding:"10px",fontSize:"16px",backgroundColor:"#007bff", color: 'white', border: 'none', cursor: 'pointer'}}>
                    Create Account
                </button>
            </form>
        </div>
    )
} 