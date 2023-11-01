import { useState } from 'react';
import './styles.css';
import { useMutation, gql } from '@apollo/client';

const defaultInput = {
  name: '',
  time: '',
  duration: 0,
  skills: '',
  guestType: 'Interviewee',
  difficulty: 'Easy-Medium',
  note: '',
  booked: false,
  meetingCode: '',
  joinUrl: '',
  guest: [],
}

interface Props {
  setShowForm: (v: boolean) => void
}

const InputForm = ({ setShowForm }: Props) => {
  const [input, setInput] = useState(defaultInput);

  const [mutateFunction, {loading, error}] = useMutation(gql`
    mutation ScheduleInterview($input: InterviewInput!) {
      scheduleInterview(input: $input) {
        id,
        name,
        skills,
        duration,
        time,
        booked,
        note,
        guestType
      }
    }
  `);

  const convertISTtoMST = (inputDateTime: string) => {
    const inputDateTimeFull = inputDateTime + ":00+05:30";
    const dateIST = new Date(inputDateTimeFull);
    dateIST.setHours(dateIST.getHours() - 11);
    dateIST.setMinutes(dateIST.getMinutes() - 30);
    const year = dateIST.getFullYear();
    const month = String(dateIST.getMonth() + 1).padStart(2, '0');
    const day = String(dateIST.getDate()).padStart(2, '0');
    const hours = String(dateIST.getHours()).padStart(2, '0');
    const minutes = String(dateIST.getMinutes()).padStart(2, '0');
    const seconds = String(dateIST.getSeconds()).padStart(2, '0');
    const convertedDateTime = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}.0000000-07:00`;
    return convertedDateTime;
  }

  function handleInput(e: React.ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target;
    setInput((prev) => ({
      ...prev,
      [name]: value,
    }));
  }

  function handleDnt(e: React.ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target;

    const convertedDateTimeMST = convertISTtoMST(value);
    console.log(convertedDateTimeMST)

    setInput((prev) => ({
      ...prev,
      [name]: convertedDateTimeMST,
    }));
  }

  function handleTextAreaInput(e: React.ChangeEvent<HTMLTextAreaElement>) {
    const { name, value } = e.target;
    setInput((prev) => ({
      ...prev,
      [name]: value,
    }));
  }
  
  function handleDurationSelect(e: React.ChangeEvent<HTMLSelectElement>) {
    const { name, value } = e.target;
    setInput((prev) => ({
      ...prev,
      [name]: parseInt(value),
    }));
  }

  function handleRadio(e: React.ChangeEvent<HTMLInputElement>) {
    const { name, value } = e.target;
    setInput((prev) => ({
      ...prev,
      [name]: value,
    }));
  }

  if (loading) return 'Submitting...';
  if (error) return `Submission error! ${error.message}`;

  return (
    <form onSubmit={(e) => {
      e.preventDefault()
      mutateFunction({ variables: { input: input }})
      setInput(defaultInput)
      setShowForm(false)
      window.location.reload()
    }}>
      <p>Schedule Mock Interview</p>

      <div className="container one">
        <input type="text" placeholder="Your Name" className="name input" onChange={handleInput} value={input.name} name='name' />
        <input type="text" placeholder="Skills" className="skills input" onChange={handleInput} value={input.skills} name='skills' />

        <div className="input">
          You are:
          <label>
            <input type="radio" name="guestType" id="interviewee" value="Interviewee" defaultChecked={true} onChange={handleRadio} />
            Interviewee
          </label>
          <label>
            <input type="radio" name="guestType" id="interviewer" value="Interviewer" onChange={handleRadio} />
            Interviewer
          </label>
        </div>

        <input type="datetime-local" name="time" id="dnt" onChange={handleDnt} />

        <div className="form-input input">
          <label>Duration: </label>
          <select value={input.duration} onChange={handleDurationSelect} name='duration' >
            <option value="10">10 minutes</option>
            <option value="15">15 minutes</option>
            <option value="20">20 minutes</option>
            <option value="30">30 minutes</option>
            <option value="45">45 minutes</option>
            <option value="60">60 minutes</option>
            <option value="90">1.5 Hours</option>
            <option value="120">2 Hours</option>
          </select>
        </div>
      </div>

      <textarea name="note" value={input.note} id="note" cols={163} rows={2} placeholder="Note(Optional)" className="input" onChange={handleTextAreaInput}></textarea>
      <button type="submit">Done</button>
    </form>
  );
};

export default InputForm;
