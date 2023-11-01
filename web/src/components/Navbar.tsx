import { Dispatch, SetStateAction } from "react";

interface NavbarProps {
  showForm: boolean
  setShowForm: Dispatch<SetStateAction<boolean>>
}

const Navbar = ({ showForm, setShowForm }: NavbarProps) => {
  return (
    <div className='navbar'>
      <div className='nav-item' onClick={() => {
        showForm == false ? setShowForm(true) : setShowForm(false);
      }}>
        Schedule New</div>
    </div>
  )
}

export default Navbar