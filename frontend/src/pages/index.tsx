import { TopBar } from '@/components/TopBar';
import { GET_ME } from '@/operations/queries/getMe'
import { GET_ALL_PATIENTS } from '@/operations/queries/getPatients';
import { GetAllPatients } from '@/operations/queries/__generated__/GetAllPatients';
import { GetMe } from '@/operations/queries/__generated__/GetMe'
import { useQuery } from '@apollo/client'
import { Alert, CircularProgress, Container, Divider, List, ListItem, ListItemIcon, ListItemText } from '@mui/material';
import { HealthAndSafety } from "@mui/icons-material";


export default function Home() {
  const {
    loading: isSelfLoading,
    data: selfData,
    error: selfError
  } = useQuery<GetMe>(GET_ME);

  const {
    loading: isPatientsLoading,
    data: patientsData,
    error: patientsError
  } = useQuery<GetAllPatients>(GET_ALL_PATIENTS);

  if (isSelfLoading || isPatientsLoading) {
    return (
      <div style={{display: 'flex', justifyContent: 'center'}}>
        <CircularProgress size={100} />
      </div>
    )
  }

  if (selfError || patientsError) {
    return (
      <div style={{display: 'flex', justifyContent: 'center'}}>
        <Alert severity='error'>Error Loading Data: {selfError ? selfError.message: patientsError?.message}</Alert>
      </div>
    )
  }

  return (
    <div>
      <TopBar name={selfData?.me.name!}></TopBar>
      <Container>
        <List>
          {patientsData?.patients.map((patient) => (
            <>
              <ListItem disablePadding>
                <ListItemIcon>
                  <HealthAndSafety />
                </ListItemIcon>
                <ListItemText primary={patient.name} secondary={patient.scans.length + " scan"}></ListItemText>
              </ListItem>
              <Divider></Divider>
            </>
          ))}
        </List>
      </Container>
    </div>
  )
}
