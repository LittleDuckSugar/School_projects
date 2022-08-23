import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import axios from 'axios';
import clsx from "clsx";

import Headers from '../Navbar/Headers';

import {
    Button,
    Box,
    Card,
    Divider,
    Grid,
    Paper,
    Typography,
    TextField,
} from "@material-ui/core";

import { makeStyles } from '@material-ui/core/styles';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundColor: 'var(--lightgray-color)',
        position: 'absolute',
        display: 'flex',
        height: '100%',
        overflow: 'hidden',
        width: '100%',
    },
    wrapper: {
        display: 'flex',
        flex: '1 1 auto',
        overflow: 'hidden',
        paddingTop: 64,
        [theme.breakpoints.up('md')]: {
            paddingLeft: 256,
        },
        [theme.breakpoints.up('lg')]: {
            paddingLeft: 256,
        },
    },
    contentContainer: {
        display: 'flex',
        flex: '1 1 auto',
        overflow: 'hidden',
    },
    content: {
        flex: '1 1 auto',
        height: '100%',
        overflow: 'auto',
    },
    pageContainer: {
        minHeight: "100%",
        marginTop: "0px",
    },
    secondRoot: {
        margin: "20px 20px",
        marginBottom: "50px",
        padding: "20px 20px",
    },
    mainTable: {
        margin: "20px 0 20px",
    },
    title: {
        color: "var(--blue-color)",
        textAlign: "center",
        margin: "30px",
    },
    divider: {
        backgroundColor: "var(--blue-color)",
        height: "5px",
        margin: "10px 0",
    },
    editForm: {
        display: "grid",
        padding: "2% 5%",
    },
    inputField: {
        margin: "10px 0",
    },
    inputCont: {
        [theme.breakpoints.down("xs")]: {
            display: "block",
        },
    },
    btnContainer: {
        textAlign: "right",
        marginTop: '20px',
        [theme.breakpoints.down("xs")]: {
            textAlign: "center",
            display: "flex",
            justifyContent: "center",
        },
    },
    button: {
        background: "var(--blue-color)",
        color: "white",
        marginLeft: "30px",
        width: "100px",
        "&:hover": {
            backgroundColor: "#113e78ec",
        },
        [theme.breakpoints.down("xs")]: {
            margin: "15px",
        },
    },
}));

const CreateCustomer = ({ className, staticContext, ...rest }) => {
    const [customer, setCustomer] = useState({});
    const history = useHistory();
    const classes = useStyles();

    const submitCustomer = async (e) => {
        console.log(customer);
        e.preventDefault();
        const url = 'http://90.22.250.124:8080/user';
        try {
            await axios.post(url, customer);
            toasterSucc();
        } catch (err) {
            err && toasterErr(err);
        }
    };

    const toasterSucc = () => {
        return (
            toast.success('Customer successfully created!', {
                position: "bottom-center",
                autoClose: 3000,
                onClose: () => history.push(`/admin/user`),
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined
            })
        );
    };

    const toasterErr = (error) => {
        return (
            toast.error(`${error}`, {
                position: "bottom-center",
                autoClose: 3000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined
            })
        );
    };

    return (
        <div className={classes.root}>
            <div className={classes.wrapper}>
                <div className={classes.contentContainer}>
                    <div className={classes.content}>

                        <Headers />

                        <div className={clsx(classes.pageContainer, className)}>
                            <Card className={clsx(classes.secondRoot, className)} elevation={2} {...rest}>
                                <Typography className={classes.title} gutterBottom variant="h4" component="h1">
                                    Create a new customer
                                </Typography>

                                <Divider className={classes.divider} />

                                <Grid>
                                    <Paper className={classes.mainTable}>
                                        <form onSubmit={(e) => submitCustomer(e)} autoComplete="off" className={classes.editForm}>
                                            <TextField
                                                className={classes.inputField}
                                                label="Firstname"
                                                type="text"
                                                name="firstname"
                                                placeholder="Enter a firstname"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{ shrink: true }}
                                                onChange={(e) => setCustomer({ ...customer, firstname: e.target.value })}
                                            />

                                            <TextField
                                                className={classes.inputField}
                                                label="Lastname"
                                                type="text"
                                                name="lastname"
                                                placeholder="Enter a lastname"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{ shrink: true }}
                                                onChange={(e) => setCustomer({ ...customer, lastname: e.target.value })}
                                            />

                                            <TextField
                                                className={classes.inputField}
                                                label="Email"
                                                type="text"
                                                name="email"
                                                placeholder="Enter an email"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{ shrink: true }}
                                                onChange={(e) => setCustomer({ ...customer, email: e.target.value })}
                                            />

                                            <Grid className={classes.inputCont} container justifyContent="space-between">
                                                <Box flex={1} mr={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Username"
                                                        type="text"
                                                        name="username"
                                                        placeholder="Enter a username"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={(e) => setCustomer({ ...customer, username: e.target.value })}
                                                    />
                                                </Box>

                                                <Box flex={1} ml={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Password"
                                                        type="password"
                                                        name="password"
                                                        placeholder="Enter a password"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={(e) => setCustomer({ ...customer, password: e.target.value })}
                                                    />
                                                </Box>
                                            </Grid>
                                            <div className={classes.btnContainer}>
                                                <Button className={classes.button} href='/admin/user'>Back</Button>
                                                <Button className={classes.button} type="submit">Save</Button>
                                            </div>
                                        </form>
                                    </Paper>
                                </Grid>
                            </Card>
                            <ToastContainer />
                        </div>
                    </div >
                </div >
            </div >
        </div >
    );
}

export default CreateCustomer;
