import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import {  NavLink } from "react-router-dom";
import { ListItemIcon, ListItemText } from '@material-ui/core';
import MailIcon from '@material-ui/icons/Mail';

const drawerWidth = 220;

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
  },
  appBar: {
    width: `calc(100% - ${drawerWidth}px)`,
    marginLeft: drawerWidth,
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  toolbar: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.default,
    padding: theme.spacing(3),
  }
}));


export const Sidebar = () => {
  const classes = useStyles();

  return (
      <Drawer
        className={classes.drawer}
        variant="permanent"
        classes={{
          paper: classes.drawerPaper,
        }}
        anchor="left"
      >
        <div style={{margin: '16px', color: 'darkcyan', fontSize: 22, fontWeight: 'bold'}}>
          Komodor
        </div>
        <List style={{padding: 0}}>
          <ListItem key="releases" component={NavLink} to="/releases" className={classes.listItem} style={{paddingTop: 4, paddingBottom: 4}}>
              <ListItemIcon style={{minWidth: 36}}><MailIcon /></ListItemIcon>
              <ListItemText primary="Releases" style={{fontSize: 14}} />
          </ListItem>
          <ListItem button key="history"  component={NavLink} to="/history" className={classes.listItem} style={{paddingTop: 4, paddingBottom: 4}}>
              <ListItemIcon style={{minWidth: 36}}><MailIcon /></ListItemIcon>
              <ListItemText primary="History" style={{fontSize: 14}} />
          </ListItem>
          <ListItem button key="pipelines" component={NavLink} to="/pipelines" className={classes.listItem} style={{paddingTop: 4, paddingBottom: 4}}>
              <ListItemIcon style={{minWidth: 36}}><MailIcon /></ListItemIcon>
              <ListItemText primary="Pipelines" style={{fontSize: 14}} />
          </ListItem>
       </List>
      </Drawer>
  );
}
