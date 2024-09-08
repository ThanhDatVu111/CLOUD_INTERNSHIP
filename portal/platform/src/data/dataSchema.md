# Data Schema
This shows data needed from the API calls and desired data format

See the following files for JSON format:
- [sampleUsers.json](./sampleUsers.json)
- [sampleProjects.json](./sampleProjects.json)
- [sampleReports.json](./sampleReports.json)


## Users

ID:             Number assigned to user by system
Name:           First and last name
Password:       Password
Mail:           Registered email
Phone:          Registered phone number
PFP:            Link to user profile picture

Title:          Job title
Department:     Team department
Status:         Indicates online, offline, away, and discharged

Calendar:       *Each list item is an object that contains an event*
                Date:           Event date
                Time:           Event time
                Description:    Brief description of event

Owned proj:     Indicates ID of a lab owned by user;
                used to reduce load times but may need
                to be refreshed
Contri. proj:   Indicates ID of any lab user edited;
                used to reduce load times but may need
                to be refreshed

Owned reports:  Indicates ID of a report owned by user;
                used to reduce load times but may need
                to be refreshed
Contri. rep.:   Indicates ID of any report user edited;
                used to reduce load times but may need
                to be refreshed


## Projects

ID:             Number assigned to project by system
Title:          Project title
Description:    Project description and overall notes
Date:           Date project created
Last modified:  Date of last modification
Last mod. by:   ID of user who made last modification
Thumbnail:      URL to thumbnail image
Owner:          ID of user who owns project
Contributors:   IDs of all contributors to the project excluding owner
Labs:           *Each array item is an object that contains a lab*
                Title:          Lab title
                Description:    Brief description of log contents
                Thumbnail:      URL to lab thumbnail image
                Date:           Date lab created
                Function:       URL to lab function code? TBD
                Locked:         Indicates whether lab can be edited
                                or not
                Contributors:   IDs of all users who contributed to
                                the log including owner
                Runs:           *Each array item is an object that*
                                *contains a "run" of the function*
                                Author:    ID of user who triggered run
                                Date:      Date of run
                                Status:    Indicates "success", "failed",
                                           "pending" based on run results
                                Results:   URL to a file of results? TBD


## Reports

ID:             Number assigned to report by system
Title:          Report title
Description:    Report description and overall notes
Date:           Date report created
Last modified:  Date of last modification
Last mod. by:   ID of user who made last modification
Thumbnail:      URL to thumbnail image
Owner:          ID of user who owns report
Contributors:   IDs of all contributors to the report excluding owner
Logs:           *Each array item is an object that contains a log*
                Title:          Log title
                Description:    Brief description of log contents
                Comment:        URL to markdown log with all log
                                detail
                Date:           Date log created
                Last modified:  Date of last modification
                Last mod. by:   ID of user who made last
                                modification
                Locked:         Indicates whether log can be edited
                                or not
                Contributors:   IDs of all users who contributed to
                                the log including owner
                Tags:           List of all tags applied to log
                Files:          List of all files uploaded along
                                with the log
                LinkedPL:       LinkedProjectLabs contains an array
                                of objects that reference labs via
                                project object id and the index of
                                actual lab object in the array