#!/bin/bash

# Ask for name
read -p "Name? (This will go on the LICENSE)
=> " name

# Ask for email
read -p "Email?
=> " email

# Ask for username
read -p "Username? (https://github.com/<username>)
=> " username

# Ask for repository
read -p "Repository? (https://github.com/$username/<repo>
=> " repository

# Ask for project name
read -p "Project name?
=> " proj_name

# Ask for project description
read -p "Short description?
=> " proj_short_desc

read -p "Long description?
=> " proj_long_desc

# Ask for docs url
read -p "Documentation URL?
=> " docs_url

# ===== Log ===== #
echo
echo
echo "===== Log ====="
echo "Name: $name"
echo "Email: $email"
echo "Username: $username"
echo "Repository: $repository"
echo "Project name: $proj_name"
echo "Project short description: $proj_short_desc"
echo "Project long description: $proj_long_desc"
echo "Docs URL: $docs_url"
echo "================"

# Ask for confirmation
while true; do
	read -p "Confirm? (y/n)
=> " confirm
	case $confirm in
	[Yy]*) break ;;
	[Nn]*) exit ;;
	*) echo "Please answer yes or no." ;;
	esac
done

# Write files
echo
echo
echo "Writing files..."

# Writing general stuff
find ./ -type f \( -iname CODEOWNERS -o -iname CITATION.cff -o -iname \*.md \) -print0 | xargs -0 sed -i -e "s/{{REPOSITORY}}/$username\/$repository/g" \
	-e "s/{{PROJECT_NAME}}/$proj_name/g" \
	-e "s/{{PROJECT_SHORT_DESCRIPTION}}/$proj_short_desc/g" \
	-e "s/{{PROJECT_LONG_DESCRIPTION}}/$proj_long_desc/g" \
	-e "s/{{DOCS_URL}}/$docs_url/g" \
	-e "s/assignees: caffeine-addictt/assignees: $username/g" \
	-e "s/contact@ngjx.org/$email/g"

# Write License
sed -i -e "s/Jun Xiang/$name/g" ./LICENSE

# Write CODEOWNERS
sed -i -e "s/caffeine-addictt/$username/g" .github/CODEOWNERS

# Write README
sed -i -e "s/Alex/$name/g" README.md

# Optional keep up-to-date
read -p "Would you like to keep up-to-date with the template? (y/n)
=> " up_to_date

case $up_to_date in
[Yy]*) {
	echo "Writing ignore file..."
	echo ".github/ISSUE_TEMPLATE/*
.github/CODEOWNERS
.github/CODESTYLE.md
.github/PULL_REQUEST_TEMPLATE.md
.github/SECURITY.md
CITATION.cff
LICENSE
README.md" >>.templatesyncignore
	echo "You can view more configuration here: https://github.com/AndreasAugustin/actions-template-sync"
} ;;
*) {
	echo "Removing syncing workflow..."
	rm .github/workflows/sync-template.yml
} ;;
esac

read -p "Would you like to keep this setup script? (y/n)
=> " keep_script

case $keep_script in
[Yy]*) {
	echo "Removing setup script..."
	rm ./setup.sh
} ;;
*) echo "Okay." ;;
esac

echo
echo "Done!
If you encounter any issues, please report it here: https://github.com/caffeine-addictt/template/issues/new?assignees=caffeine-addictt&labels=Type%3A+Bug&projects=&template=1-bug-report.md&title=[Bug]+"
