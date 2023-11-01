import { useState } from "react"
import { TypeInterview } from "../schema"
import { humanTime } from "../utils"
import './styles.css'
import { useMutation, gql } from "@apollo/client";

const defaultState = {
  name: "",
  email: "",
}

const Interview = (props: { interview: TypeInterview }) => {
  const { interview } = props
  const [showDetails, setShowDetails] = useState(false)
  const [guest, setGuest] = useState(defaultState);
  const [mutateFunction] = useMutation(gql`
    mutation bookmeeting($input: BookInterview!) {
      bookInterview(input: $input) {
        id
      }
    }
  `)

  function handleGuestInput(e: React.ChangeEvent<HTMLInputElement>) {
    const {name, value} = e.target;
    setGuest(prev => ({
      ...prev,
      [name]: value,
    }))
  }

  function handleBooking(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()
    mutateFunction({ variables: { input: {
      interviewId: interview.id,
      name: guest.name,
      email: guest.email
    } } })
    //setGuest(defaultState)
    //window.location.reload()
  }

  let guestList = "";
  if (interview.guest.length > 0) {
    let count = 0;
    interview.guest.forEach(gt => {
      if (count > 0) {
        guestList += `, ${gt.name}`
      } else {
        guestList += gt.name
      }
      count++
    })
  }
  
  return (
    <div className="card" key={interview.id}>
      <div className="compact" onClick={() => {
        showDetails == false ? setShowDetails(true) : setShowDetails(false);
      }}>
        <div className="items skills">{`${interview.skills}`}</div>
        <div className="items dnt">{`${humanTime(interview.time)}`}</div>
        <div className="items duration">{`${interview.duration}m`}</div>
        <div className="items bnk">{interview.booked ? "Booked" : "Not Booked"}</div>
        <div className="items name">{`${interview.name}(${interview.guestType})`}</div>
        <div className="items note">{interview.note}</div>
      </div>

      {showDetails && (
        <div className="detailed">
          <div className="info">
            <div className="note">Note: {interview.note}</div>
            {interview.booked && (
              <div className="guest">{`Guest: ${guestList}`}</div>
            )}
          </div>
          <form onSubmit={handleBooking}>
              <input type="text" id="name" name="name" placeholder="Guest Name" onChange={handleGuestInput} />
              <input type="email" id="email" name="email" placeholder="Email..." onChange={handleGuestInput} />
              <button type="submit">Confirm Booking</button>
          </form>
        </div>
      )}
    </div>
  )
}

export default Interview