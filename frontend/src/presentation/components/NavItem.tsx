import { Link } from "react-router-dom";

export const NavItem: React.FC<{ to: string; label: string }> = ({ to, label }) => (
  <Link to={to} className='hover:underline transition duration-200 ease-in-out'>
    {label}
  </Link>
);