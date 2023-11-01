import Interview from "../components/Interview"
import InputForm from "../components/InputForm";
import { TypeInterview } from "../schema";
import { gql, useQuery } from '@apollo/client';
import './style.css'
import Navbar from "../components/Navbar";
import { useState } from "react";

const Interviews = () => {
  const [showForm, setShowForm] = useState(false)
  const { loading, error, data } = useQuery(
    gql`
      query interviews {
        allInterviews {
          id
          name
          duration
          skills
          time
          difficulty
          booked
          note
          guestType
          guest {
            name
            email
          }
        }
      }`
  );

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

  function provideUniqueKey() {
    let key = 0;
    return function() {
      return key++
    }
  }
  const getKey = provideUniqueKey();

  return (
      <>
        <Navbar showForm={showForm} setShowForm={setShowForm} />
        {showForm ? <InputForm setShowForm={setShowForm} /> : null}
        <div className="title">
          <p className="skills">Skills</p>
          <p className="dnt">Date & Time</p>
          <p className="duration">Duration</p>
          <p className="status">Status</p>
          <p className="name">Name</p>
          <p className="note">Note...</p>
        </div>
        {data && data.allInterviews.map((interview: TypeInterview) => {
          return <Interview key={getKey()} interview={interview} />;
        })}
      </>
  )
}

export default Interviews