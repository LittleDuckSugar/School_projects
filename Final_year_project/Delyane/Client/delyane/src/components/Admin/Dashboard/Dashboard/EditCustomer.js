import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import axios from 'axios';
import clsx from 'clsx';

import Headers from '../Navbar/Headers';

import {
    Button,
    Card,
    Divider,
    Grid,
    Paper,
    TextField,
    Typography,
} from '@material-ui/core';

import DeleteForeverRoundedIcon from '@material-ui/icons/DeleteForeverRounded';

import { makeStyles } from '@material-ui/core';

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
        minHeight: '100%',
        marginTop: '0px',
    },
    secondRoot: {
        margin: '20px 20px',
        marginBottom: '50px',
        padding: '20px 20px',
    },
    mainTable: {
        margin: '20px 0 20px',
    },
    title: {
        color: 'var(--blue-color)',
        textAlign: 'center',
        margin: '30px',
    },
    divider: {
        backgroundColor: 'var(--blue-color)',
        height: '5px',
        margin: '10px 0',
    },
    editForm: {
        display: 'grid',
        padding: '2% 5%',
    },
    inputField: {
        margin: '10px 0',
    },
    inputCont: {
        [theme.breakpoints.down('xs')]: {
            display: 'block',
        },
    },
    btnContainer: {
        textAlign: 'right',
        marginTop: '20px',
        [theme.breakpoints.down('xs')]: {
            textAlign: 'center',
            display: 'flex',
            justifyContent: 'center',
        },
    },
    button: {
        background: 'var(--blue-color)',
        color: 'white',
        marginLeft: '30px',
        width: '100px',
        '&:hover': {
            backgroundColor: '#113e78ec',
        },
        [theme.breakpoints.down('xs')]: {
            margin: '15px',
        },
    },
    deleteBtn: {
        background: 'red',
        color: 'white',
        minInlineSize: 'auto',
        float: 'right',
        '&:hover': {
            backgroundColor: 'var(--pwr-btn-alarm)',
        },
    },
}));


const EditCustomer = ({ className, staticContext, ...rest }) => {
    const [saved, setSaved] = useState(false);
    const { uuid } = useParams();
    const history = useHistory();
    const classes = useStyles();

    const [customer, setCustomer] = useState({
        firstname: '',
        lastname: '',
        email: '',
        username: ''
    });

    const [updatedCustomer, setUpdatedCustomer] = useState({
        firstname: '',
        lastname: '',
        email: '',
        username: ''
    })

    const handleUpdateCustomer = (e) => {
        setUpdatedCustomer({
            ...updatedCustomer,
            [e.target.name]: e.target.value
        })
    }

    const toasterSucc = () => {
        return (
            toast.success('Customer successfully edited/deleted!', {
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

    useEffect(() => {
        const getDatas = async () => {
            try {
                const result = await axios.get(`http://90.22.250.124:8080/user/${uuid}`)
                setCustomer(result.data);
                setUpdatedCustomer(result.data);
            } catch (err) {
                console.log(err)
            }
        };
        getDatas();

        return () => setCustomer([]);

    }, [uuid]);

    const updateCustomer = async (e) => {
        const url = `http://90.22.250.124:8080/user/${uuid}`;
        try {
            await axios.put(url, updatedCustomer).then((res) => console.log(res))
            toasterSucc();
            setSaved(true);
        } catch (err) {
            err && toasterErr(err);
        }
    };

    const deleteCustomer = async () => {
        try {
            await axios.delete(`http://90.22.250.124:8080/user/${uuid}`);
            toasterSucc();
        } catch (err) {
            err && toasterErr(err);
        }
    }

    console.log(customer)

    return customer ? (
        <div className={classes.root}>
            <div className={classes.wrapper}>
                <div className={classes.contentContainer}>
                    <div className={classes.content}>

                        <Headers />
                        <div className={clsx(classes.pageContainer, className)}>
                            <Card className={clsx(classes.secondRoot, className)} elevation={2} {...rest}>
                                <Button className={classes.deleteBtn} onClick={deleteCustomer}>
                                    <DeleteForeverRoundedIcon />
                                </Button>
                                <Typography className={classes.title} gutterBottom variant="h4" component="h1">
                                    Edit a customer
                                </Typography>

                                <Divider className={classes.divider} />

                                <Grid>
                                    <Paper className={classes.mainTable}>
                                        <form className={classes.editForm} onSubmit={(e) => updateCustomer(e)} autoComplete="off" >
                                            <TextField
                                                className={classes.inputField}
                                                label="Firstname"
                                                type="text"
                                                name="firstname"
                                                placeholder="Enter a firstname"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{
                                                    shrink: true,
                                                }}
                                                value={updatedCustomer.firstname || ''}
                                                onChange={handleUpdateCustomer}
                                            />

                                            <TextField
                                                className={classes.inputField}
                                                label="Lastname"
                                                type="text"
                                                name="lastname"
                                                placeholder="Enter a lastname"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{
                                                    shrink: true,
                                                }}
                                                value={updatedCustomer.lastname || ''}
                                                onChange={handleUpdateCustomer}
                                            />

                                            <TextField
                                                className={classes.inputField}
                                                label="Email"
                                                type="text"
                                                name="email"
                                                placeholder="Enter an email"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{
                                                    shrink: true,
                                                }}
                                                value={updatedCustomer.email || ''}
                                                onChange={handleUpdateCustomer}
                                            />

                                            <TextField
                                                className={classes.inputField}
                                                label="Username"
                                                type="text"
                                                name="username"
                                                placeholder="Enter a username"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{
                                                    shrink: true,
                                                }}
                                                value={updatedCustomer.username || ''}
                                                onChange={handleUpdateCustomer}
                                            />

                                            <div className={classes.btnContainer}>
                                                {customer === updatedCustomer || saved ? (
                                                    <Button href='/admin/user' className={classes.button}>Back</Button>
                                                ) : (
                                                    <Button href='/admin/user' className={classes.button}>Cancel</Button>
                                                )}
                                                <Button className={classes.button} type="submit">Save</Button>
                                            </div>
                                        </form>
                                    </Paper>
                                </Grid>
                            </Card>
                            <ToastContainer />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    ) : null;
}

export default EditCustomer;
