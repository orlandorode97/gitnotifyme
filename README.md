# Gitnotifyme

Gitnotifyme is a tool that listens to github events through a webhook for pull requests

Note: Only works for a single user. I need to find a time to extend this functionally

# Get started
1. Install in your local environment this webhook payload delivery https://smee.io/
2. Run the following command: `smee -u https://smee.io/4OWRIU80tZ73fo -t http://localhost:3000/events`
3. At the root of the folder execute `make build` to create an executable file.
4. Run `./gitnotifyme`
5. Recieve notifications for PRs events.

# Demo
![demo4](https://github.com/orlandorode97/gitnotifyme/assets/34588445/65aaf061-98ce-461d-bcbd-eba1705ca63c)
