import React, { useState } from 'react';
import clsx from "clsx";
import axios from 'axios';

import Headers from '../Navbar/Headers';

import {
    Avatar,
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
    imgContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        [theme.breakpoints.down('xs')]: {
            display: 'flex',
            flexDirection: 'column',
            textAlign: 'center',
        },
    },
    avatar: {
        margin: '20px',
        cursor: 'pointer',
        borderRadius: 15,
        width: '180px',
        height: '130px',
        border: '2px solid var(--gmi-color)',
        [theme.breakpoints.down('xs')]: {
            marginTop: '15px',
        },
    },
    imgNota: {
        color: 'var(--pwr-btn-alarm)',
        textAlign: 'center',
        margin: '0 auto 30px',
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

const CreatePainting = ({ className, staticContext, ...rest }) => {
    const [title, setTitle] = useState('');
    const [price, setPrice] = useState('');
    const [technical, setTechnical] = useState('');
    const [dimension, setDimension] = useState('');
    const [authentification, setAuthentification] = useState('');
    const [support, setSupport] = useState('');
    const [description, setDescription] = useState('');
    const [chosePicture, setChosePicture] = useState('');
    const classes = useStyles();
    // const history = useHistory();

    const handleChangeTitle = (e) => {
        setTitle(e.target.value);
    };

    const handleChangePrice = (e) => {
        setPrice(e.target.value);
    };

    const handleChangeTechnical = (e) => {
        setTechnical(e.target.value);
    };

    const handleChangeDimension = (e) => {
        setDimension(e.target.value);
    };

    const handleChangeAuthentification = (e) => {
        setAuthentification(e.target.value);
    };

    const handleChangeSupport = (e) => {
        setSupport(e.target.value);
    };

    const handleChangeDescription = (e) => {
        setDescription(e.target.value);
    };

    const submitPainting = async (e) => {
        e.preventDefault();
        const formData = new FormData();
        formData.append('title', title);
        formData.append('price', price);
        formData.append('technical', technical);
        formData.append('dimension', dimension);
        formData.append('authentification', authentification);
        formData.append('support', support);
        formData.append('description', description);
        formData.append('avatar', chosePicture);
        try {
            await axios.post('http://90.22.250.124:8080/product', formData);
        } catch (err) {
            console.log(err)
        }
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
                                        <form onSubmit={(e) => submitPainting(e)} encType="multipart/form-data" method="post" autoComplete="off" className={classes.editForm}>

                                            <div className={classes.imgContainer}>
                                                {chosePicture ? (
                                                    <Avatar
                                                        className={classes.avatar}
                                                        src={URL.createObjectURL(chosePicture)}
                                                        alt={``}
                                                    />
                                                ) : (
                                                    <Avatar
                                                        className={classes.avatar}
                                                        src={``}
                                                        alt={'No picture'}
                                                    />
                                                )}
                                                <input
                                                    className={classes.btnfile}
                                                    type="file"
                                                    id="file"
                                                    name="avatar"
                                                    onChange={(e) => {
                                                        setChosePicture(e.target.files[0]);
                                                    }}
                                                />
                                            </div>

                                            <Grid className={classes.inputCont} container justifyContent="space-between">

                                                <Box flex={2} mr={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Title"
                                                        type="text"
                                                        name="title"
                                                        placeholder="Enter a title"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={handleChangeTitle}
                                                    />
                                                </Box>

                                                <Box flex={1} ml={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Price"
                                                        type="number"
                                                        name="price"
                                                        placeholder="Enter a price"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={handleChangePrice}
                                                    />
                                                </Box>
                                            </Grid>

                                            <Grid className={classes.inputCont} container justifyContent="space-between">
                                                <Box flex={1} mr={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Technical"
                                                        type="text"
                                                        name="technical"
                                                        placeholder="Enter a technical"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={handleChangeTechnical}
                                                    />
                                                </Box>

                                                <Box flex={1} ml={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Dimension"
                                                        type="text"
                                                        name="dimension"
                                                        placeholder="Enter a dimension"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={handleChangeDimension}
                                                    />
                                                </Box>
                                            </Grid>


                                            <Grid className={classes.inputCont} container justifyContent="space-between">
                                                <Box flex={1} mr={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Authentification"
                                                        type="text"
                                                        name="authentification"
                                                        placeholder="Enter an authentification"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={handleChangeAuthentification}
                                                    />
                                                </Box>

                                                <Box flex={1} ml={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Support"
                                                        type="text"
                                                        name="support"
                                                        placeholder="Enter a support"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                        onChange={handleChangeSupport}
                                                    />
                                                </Box>
                                            </Grid>

                                            <TextField
                                                className={classes.inputField}
                                                label="Description"
                                                type="text"
                                                name="description"
                                                placeholder="Enter a description"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{ shrink: true }}
                                                multiline
                                                minRows={4}
                                                onChange={handleChangeDescription}
                                            />

                                            <div className={classes.btnContainer}>
                                                <Button className={classes.button} href='/admin/painting'>Back</Button>
                                                <Button className={classes.button} type="submit">Save</Button>
                                            </div>
                                        </form>
                                    </Paper>
                                </Grid>
                            </Card>
                            {/* <ToastContainer /> */}
                        </div>
                    </div >
                </div >
            </div >
        </div >
    );
}

export default CreatePainting;
