
# Artist Filters Project

This project allows users to filter artists based on various criteria, including the number of members, creation date, first album date, and location. The application uses an interactive UI built with HTML and Tailwind CSS to enable users to adjust filter settings dynamically and see the results in real-time.

## Features

- **Filter by Number of Members**: Use a sliding range to filter artists by the number of members in their group.
- **Filter by Creation Date**: Enter a range to filter artists by their creation date.
- **Filter by First Album Date**: Choose a date range to filter artists based on when their first album was released.
- **Filter by Location**: Search for artists based on their location, with suggestions shown as the user types.
- **Dynamic Results**: The artist results update instantly based on the selected filters.

## Getting Started

To run the project locally, follow the steps below.

### Prerequisites

- A modern web browser (Chrome, Firefox, etc.).
- A basic text editor or IDE (VS Code, Sublime Text, etc.).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/artist-filters.git
   ```
   
2. Navigate to the project directory:
   ```bash
   cd artist-filters
   ```

3. Open `index.html` in your preferred web browser.

### Usage

- Use the **Number of Members** slider to select a value. Artists with fewer than or equal to the selected number of members will be shown.
- Enter a range of years in the **Creation Date** fields to filter artists based on their formation year.
- Use the **First Album Date** fields to filter artists based on the release date of their first album.
- Search for specific **Locations** by typing into the location search box. Location suggestions will appear as you type, and selecting one will filter the artists by that location.
- The results will be displayed dynamically based on your filter selections.

### Example

If you want to filter artists with fewer than 10 members and located in the "USA", adjust the slider to "10" for members and type "USA" in the location search field. The displayed results will show only those artists who meet both criteria.

### Technologies Used

- **HTML**: For structuring the page.
- **CSS (Tailwind CSS)**: For styling the page and creating a responsive layout.
- **JavaScript**: To handle dynamic filtering and updating the displayed results based on user input.

### Customization

You can easily extend this project by adding more filters or modifying the existing ones. For example, you could add a filter for genres or music styles. Additionally, the artist data can be replaced with a backend API to fetch real-time information.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

